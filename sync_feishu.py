#!/usr/bin/env python3
"""
sync_feishu.py - 同步飞书知识库文章到 110claw 数据库

用法:
    python3 sync_feishu.py              # 执行一次同步
    python3 sync_feishu.py --schedule   # 后台定时同步（每 6 小时）

环境变量 (或从 backend/.env 读取):
    DB_DSN      - MySQL 连接串
    WIKI_TOKEN  - 飞书 wiki token (默认: OCGUw7De8i1zFUkpWPQc6oi8nOb)
"""

import argparse
import json
import os
import signal
import sys
import time
import urllib.request
import urllib.error
from datetime import datetime

# Playwright 会自动处理 Chromium
from playwright.sync_api import sync_playwright

# ──────────────────────────────────────────────────────
# 配置
# ──────────────────────────────────────────────────────

BASE_URL = "https://waytoagi.feishu.cn"
DEFAULT_WIKI_TOKEN = "OCGUw7De8i1zFUkpWPQc6oi8nOb"

# 从 .env 文件读取配置
def load_env():
    env_path = os.path.join(os.path.dirname(__file__), "backend", ".env")
    if os.path.exists(env_path):
        with open(env_path) as f:
            for line in f:
                line = line.strip()
                if not line or line.startswith("#"):
                    continue
                if "=" in line:
                    key, val = line.split("=", 1)
                    # 只设置未定义的环境变量
                    if key.strip() not in os.environ:
                        os.environ[key.strip()] = val.strip().strip('"').strip("'")

# ──────────────────────────────────────────────────────
# 数据库操作
# ──────────────────────────────────────────────────────

def get_db_connection():
    """获取 MySQL 连接 (使用 PyMySQL)"""
    try:
        import pymysql
    except ImportError:
        print("安装 PyMySQL: pip3 install PyMySQL")
        sys.exit(1)

    dsn = os.environ.get("DB_DSN", "")
    # 解析 DSN: user:pass@tcp(host:port)/dbname?options
    import re
    m = re.match(r"([^:]+):([^@]+)@tcp\(([^:]+):(\d+)\)/([^?]+)", dsn)
    if not m:
        print(f"无法解析 DB_DSN: {dsn}")
        sys.exit(1)

    user, password, host, port, database = m.groups()
    return pymysql.connect(
        host=host, port=int(port), user=user, password=password, database=database,
        charset="utf8mb4", cursorclass=pymysql.cursors.DictCursor
    )

def upsert_article(conn, title, description, content, source_url, sort_order):
    """插入或更新文章，返回状态: 'inserted', 'updated', 'skipped'"""
    cursor = conn.cursor()
    # 先检查是否存在及其当前值
    cursor.execute(
        "SELECT id, title, description, sort_order FROM learn_steps WHERE source_url = %s",
        (source_url,)
    )
    row = cursor.fetchone()

    if row:
        # 检查是否有变化
        if (row['title'] == title and
            row['description'] == description and
            row['sort_order'] == sort_order):
            # 无变化，跳过
            return 'skipped'
        # 有变化，更新
        cursor.execute("""
            UPDATE learn_steps
            SET title = %s, description = %s, content = %s, sort_order = %s, updated_at = NOW()
            WHERE source_url = %s
        """, (title, description, content, sort_order, source_url))
        conn.commit()
        return 'updated'
    else:
        # 插入新记录
        cursor.execute("""
            INSERT INTO learn_steps (title, description, content, source_url, sort_order, status, created_at, updated_at)
            VALUES (%s, %s, %s, %s, %s, 1, NOW(), NOW())
        """, (title, description, content, source_url, sort_order))
        conn.commit()
        return 'inserted'

def mark_deleted(conn, existing_urls):
    """标记不在列表中的文章为已删除"""
    if not existing_urls:
        return 0
    cursor = conn.cursor()
    placeholders = ",".join(["%s"] * len(existing_urls))
    cursor.execute(f"""
        UPDATE learn_steps SET status = 0, updated_at = NOW()
        WHERE source_url NOT IN ({placeholders}) AND status = 1
    """, tuple(existing_urls))
    conn.commit()
    return cursor.rowcount

# ──────────────────────────────────────────────────────
# 飞书 API 爬取
# ──────────────────────────────────────────────────────

def fetch_wiki_data(wiki_token):
    """使用 Playwright 获取飞书 wiki 数据"""

    children_data = [None]

    def on_response(resp):
        if 'get_node_child_paging' in resp.url:
            try:
                children_data[0] = resp.json()
            except:
                pass

    print(f"[{datetime.now().isoformat()}] 开始抓取飞书 wiki: {wiki_token}")

    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        ctx = browser.new_context()
        page = ctx.new_page()
        page.on("response", on_response)

        try:
            page.goto(f"{BASE_URL}/wiki/{wiki_token}", wait_until="networkidle", timeout=60000)
            page.wait_for_timeout(5000)
        except Exception as e:
            print(f"页面加载失败: {e}")
        finally:
            browser.close()

    if not children_data[0]:
        print("未能获取数据")
        return []

    data = children_data[0]
    if data.get("code") != 0:
        print(f"API 错误: {data.get('msg')}")
        return []

    nodes = data.get("data", {}).get("nodes", {})
    if not nodes:
        print("没有找到文章")
        return []

    # 按 sort_id 降序排列（最新的在前）
    items = sorted(nodes.values(), key=lambda x: x.get("sort_id", 0), reverse=True)
    print(f"获取到 {len(items)} 篇文章")

    return items

def extract_description(item):
    """从文章元数据提取描述"""
    # 尝试从 detail_info 获取创建时间作为描述的一部分
    detail = item.get("detail_info", {})
    create_time = detail.get("create_time")
    if create_time:
        try:
            ts = int(create_time)
            dt = datetime.fromtimestamp(ts)
            return f"创建于 {dt.strftime('%Y-%m-%d')}"
        except:
            pass
    return "学习 OpenClaw 实战技巧"

# ──────────────────────────────────────────────────────
# 主同步逻辑
# ──────────────────────────────────────────────────────

def sync_once(wiki_token):
    """执行一次同步"""
    items = fetch_wiki_data(wiki_token)
    if not items:
        print("没有数据需要同步")
        return 0

    conn = get_db_connection()
    existing_urls = []

    # 统计计数
    stats = {'inserted': 0, 'updated': 0, 'skipped': 0, 'failed': 0}

    try:
        for i, item in enumerate(items):
            title = item.get("title", "无标题")
            wiki_token_item = item.get("wiki_token")
            source_url = item.get("url") or f"{BASE_URL}/wiki/{wiki_token_item}"
            description = extract_description(item)

            # 暂时不获取文章正文，只存元数据
            # 正文需要单独请求每个文档，耗时较长
            content = ""

            sort_order = len(items) - i  # 倒序，最新的 sort_order 最大

            try:
                status = upsert_article(conn, title, description, content, source_url, sort_order)
                existing_urls.append(source_url)
                stats[status] += 1
                # 只打印新增和更新的
                if status == 'inserted':
                    print(f"  [新增] {title[:50]}")
                elif status == 'updated':
                    print(f"  [更新] {title[:50]}")
            except Exception as e:
                stats['failed'] += 1
                print(f"  [失败] {title[:30]} - {e}")

        # 标记删除不存在于当前列表的文章
        deleted_count = mark_deleted(conn, existing_urls)
        if deleted_count:
            print(f"  [下线] {deleted_count} 篇文章")

        conn.close()
        print(f"[{datetime.now().isoformat()}] 同步完成: 新增 {stats['inserted']}, 更新 {stats['updated']}, 跳过 {stats['skipped']}, 失败 {stats['failed']}")
        return stats['inserted'] + stats['updated']

    except Exception as e:
        print(f"同步失败: {e}")
        conn.close()
        return 0

def schedule_sync(wiki_token, interval_hours=6):
    """定时同步"""
    print(f"启动定时同步，间隔 {interval_hours} 小时")

    # 优雅退出
    stop_flag = [False]
    def handler(signum, frame):
        print("\n收到退出信号...")
        stop_flag[0] = True
    signal.signal(signal.SIGINT, handler)
    signal.signal(signal.SIGTERM, handler)

    while not stop_flag[0]:
        sync_once(wiki_token)
        if stop_flag[0]:
            break
        print(f"\n下次同步时间: {interval_hours} 小时后\n")
        for _ in range(interval_hours * 3600):
            if stop_flag[0]:
                break
            time.sleep(1)

# ──────────────────────────────────────────────────────
# 入口
# ──────────────────────────────────────────────────────

def main():
    parser = argparse.ArgumentParser(description="同步飞书知识库到 110claw")
    parser.add_argument("--schedule", action="store_true", help="启动定时同步模式")
    parser.add_argument("--interval", type=int, default=6, help="定时同步间隔（小时）")
    parser.add_argument("--token", default=None, help="飞书 wiki token")
    args = parser.parse_args()

    # 加载环境变量
    load_env()

    wiki_token = args.token or os.environ.get("FEISHU_WIKI_TOKEN") or DEFAULT_WIKI_TOKEN

    if args.schedule:
        schedule_sync(wiki_token, args.interval)
    else:
        sync_once(wiki_token)

if __name__ == "__main__":
    main()