export const learningSteps = [
  {
    day: 1,
    title: '初识 110Claw',
    desc: '了解 AI 助手的基本概念，以及 110Claw 能为你做什么。',
    duration: '约 30 分钟',
    content: `# Day 1：初识 110Claw

欢迎开始你的 110Claw 学习之旅！今天我们将了解 110Claw 是什么、它能做什么，以及如何开始使用。

## 什么是 110Claw？

110Claw 是一个面向中文用户的开源 AI 助手平台，基于大语言模型构建，具有以下核心特点：

- **免费开源**：核心功能完全免费，代码完全透明
- **安全可信**：支持本地模型，数据不出设备
- **可扩展**：通过 Skill 插件系统无限扩展能力
- **中文优先**：深度优化中文对话体验

## 110Claw 能做什么？

### 基础对话
- 回答问题、写作辅助、头脑风暴
- 翻译、摘要、内容改写

### 代码助手
- 代码生成、调试、代码审查
- 技术文档编写

### 自动化
- 定时任务、工作流自动化
- 系统监控与告警

## 安装与配置

### 安装 110Claw

\`\`\`bash
# macOS / Linux
curl -fsSL https://install.110claw.com | sh

# Windows（PowerShell）
iwr https://install.110claw.com/win | iex
\`\`\`

### 首次配置

安装完成后，运行 \`claw init\` 进行初始化配置。你可以选择使用云端模型（需要 API Key）或本地模型（需要提前安装 Ollama）。

## 今日练习

1. 完成 110Claw 安装
2. 运行你的第一条指令：\`claw chat "你好，介绍一下你自己"\`
3. 在对话框中问 110Claw："你能帮我做哪些事情？"

## 小结

恭喜你完成了第一天的学习！明天我们将深入学习如何与 AI 助手进行高效对话。`,
  },
  {
    day: 2,
    title: '深入对话',
    desc: '掌握与 AI 助手对话的技巧，让沟通更高效、更自然。',
    duration: '约 45 分钟',
    content: `# Day 2：深入对话

AI 助手的能力上限，很大程度上取决于你的提问方式。今天我们来学习如何与 110Claw 进行高效对话。

## 提示词（Prompt）的艺术

### 基本原则

1. **清晰具体**：描述清楚你的需求和上下文
2. **角色设定**：告诉 AI 以什么角色来回答
3. **格式要求**：明确输出格式（列表、表格、代码等）
4. **示例引导**：提供一两个示例，帮助 AI 理解你的期望

### 对比示例

❌ 不好的提示：
> "写一篇文章"

✅ 好的提示：
> "以技术博主的口吻，为中级 Python 开发者写一篇 800 字的文章，介绍异步编程的核心概念，包含代码示例，使用 Markdown 格式"

## 多轮对话技巧

### 建立上下文

110Claw 会记住当前对话的历史，你可以循序渐进地深入话题：

\`\`\`
用户：帮我写一个 Python 函数，计算斐波那契数列
AI：[给出代码]
用户：加上递归记忆化优化
AI：[改进代码]
用户：现在加上单元测试
AI：[完整版本]
\`\`\`

### 纠正与迭代

如果 AI 的回答不符合预期，直接告诉它：
- "不对，我想要的是……"
- "这个方向很好，但能更简洁吗？"
- "换一种角度再试试"

## 系统提示（System Prompt）

110Claw 支持自定义系统提示，让 AI 以固定角色服务你：

\`\`\`bash
claw chat --system "你是一位严格的代码审查员，会指出所有潜在问题" "审查这段代码：..."
\`\`\`

## 今日练习

1. 尝试"零样本"、"少样本"两种提示方式，对比效果
2. 创建一个"代码助手"角色对话，完成一个小功能的开发
3. 练习多轮对话，逐步完善一篇技术文章的大纲

## 小结

掌握提示词技巧是使用 AI 助手的核心能力，需要反复实践。明天我们将学习如何让 AI 处理文件和代码。`,
  },
  {
    day: 3,
    title: '文件与代码',
    desc: '让 AI 助手帮你处理文件、写代码、执行任务。',
    duration: '约 60 分钟',
    content: `# Day 3：文件与代码

今天我们来学习 110Claw 最实用的功能之一——处理文件和代码。

## 文件处理

### 读取文件

110Claw 可以直接读取并分析各种文件：

\`\`\`bash
# 分析代码文件
claw chat --file main.py "这段代码有什么潜在问题？"

# 总结文档
claw chat --file report.pdf "用三句话总结这份报告的核心内容"

# 处理 CSV 数据
claw chat --file data.csv "分析这份数据，找出异常值"
\`\`\`

### 支持的文件格式

| 类型 | 格式 |
|------|------|
| 代码 | Python、Go、JavaScript、Java、C++ 等几乎所有语言 |
| 文档 | PDF、Word、Markdown、TXT |
| 数据 | CSV、JSON、XML、Excel |
| 图片 | PNG、JPG（需要视觉模型支持）|

## 代码生成与调试

### 生成代码

\`\`\`bash
claw chat "用 Go 写一个 HTTP 中间件，记录请求耗时并输出日志"
\`\`\`

### 调试代码

\`\`\`bash
# 直接粘贴错误信息
claw chat "我遇到了这个错误：[错误信息]，代码是：[代码]，如何修复？"
\`\`\`

### 代码重构

\`\`\`bash
claw chat --file legacy.py "将这段代码重构为面向对象风格，添加类型注解"
\`\`\`

## 批量文件处理

使用 110Claw CLI 的 glob 模式处理多个文件：

\`\`\`bash
# 为所有 Python 文件生成文档注释
claw batch --files "src/**/*.py" "为这个函数添加 docstring"
\`\`\`

## 今日练习

1. 找一段你工作中的代码，让 110Claw 做代码审查
2. 让 110Claw 帮你写一个简单的自动化脚本
3. 尝试让 AI 读取一份文档并生成摘要

## 小结

文件处理能力让 110Claw 真正融入你的日常工作流程。明天我们将探索 AI 助手的网络能力。`,
  },
  {
    day: 4,
    title: '网络能力',
    desc: '搜索、抓取、API 调用，让 AI 助手连接互联网。',
    duration: '约 45 分钟',
    content: `# Day 4：网络能力

今天我们将学习如何让 110Claw 连接互联网，获取实时信息并与外部服务交互。

## 网页搜索

110Claw 内置了联网搜索能力（需要开启网络权限）：

\`\`\`bash
# 搜索最新信息
claw chat --search "2026年最流行的前端框架有哪些？"

# 研究特定话题
claw chat --search "查找 Go 1.23 的主要新特性"
\`\`\`

## 网页内容抓取

使用 \`web-scraper\` Skill 抓取网页内容：

\`\`\`bash
# 安装 web-scraper
claw skill install web-scraper

# 抓取并分析网页
claw chat "抓取 https://example.com 的内容，并总结主要信息"
\`\`\`

## API 调用

110Claw 可以帮你调用外部 API 并处理结果：

\`\`\`bash
# 调用天气 API
claw chat "调用 OpenWeatherMap API 获取北京今天的天气，API Key 是 xxx"

# 处理 REST API 响应
claw chat "帮我调用这个 API 并解析响应：GET https://api.example.com/users"
\`\`\`

## 实时数据监控

结合定时任务，可以实现持续的数据监控：

\`\`\`bash
# 每小时检查网站状态
claw schedule "0 * * * *" "检查 https://mysite.com 是否正常，如果不正常发送告警"
\`\`\`

## 安全注意事项

- 网络请求会消耗流量，注意使用频率
- 不要在请求中包含敏感信息（密码、Token）
- 可在 \`~/.claw/config.yaml\` 中配置允许访问的域名白名单

## 今日练习

1. 让 110Claw 搜索并总结一个你感兴趣的技术话题
2. 安装 web-scraper Skill，抓取一个新闻网站的最新标题
3. 尝试让 AI 调用一个公开 API（如 GitHub API）

## 小结

网络能力让 AI 助手从"知识库"变成了"实时助理"。明天我们来学习如何通过 Skill 插件进一步扩展能力。`,
  },
  {
    day: 5,
    title: '技能扩展',
    desc: '安装社区插件，让 AI 助手学会更多能力。',
    duration: '约 60 分钟',
    content: `# Day 5：技能扩展

Skill 插件系统是 110Claw 最强大的特性之一，它让你可以按需为 AI 助手添加各种能力。

## Skill 插件是什么？

Skill 是 110Claw 的能力扩展单元，类似于浏览器的扩展程序。每个 Skill 封装了特定领域的功能，可以通过 ClawHub 一键安装。

## 管理 Skills

\`\`\`bash
# 搜索 Skill
claw skill search "代码审查"

# 安装 Skill
claw skill install code-reviewer

# 查看已安装的 Skill
claw skill list

# 更新所有 Skill
claw skill upgrade --all

# 卸载 Skill
claw skill remove code-reviewer
\`\`\`

## 推荐必装 Skills

### 开发类
- **code-reviewer**：自动代码审查
- **docker-manager**：Docker 容器管理
- **git-assistant**：Git 操作助手

### 效率类
- **web-scraper**：网页抓取
- **file-organizer**：文件自动整理
- **note-taker**：智能笔记管理

### 数据类
- **data-analyzer**：数据分析与可视化
- **csv-processor**：CSV 文件批量处理

## 开发自己的 Skill

\`\`\`bash
# 初始化新 Skill 项目（支持 Go/Python/TypeScript）
claw skill create my-skill --lang python

# 项目结构
# my-skill/
# ├── skill.yaml    # 元数据配置
# ├── main.py       # 入口文件
# └── README.md     # 使用说明
\`\`\`

## skill.yaml 配置示例

\`\`\`yaml
name: my-skill
version: 1.0.0
description: 我的第一个 Skill
author: your-name
category: productivity
permissions:
  - filesystem.read
  - network.http
\`\`\`

## 今日练习

1. 安装 3 个你感兴趣的 Skill 并测试
2. 尝试初始化一个自己的 Skill 项目
3. 阅读 ClawHub 上评分最高的 Skill 的源码，学习最佳实践

## 小结

Skill 生态让 110Claw 的能力几乎没有上限。明天我们来学习如何让 AI 助手自主工作——自动化！`,
  },
  {
    day: 6,
    title: '自动化',
    desc: '定时任务、心跳监测、主动提醒，让 AI 助手自主工作。',
    duration: '约 60 分钟',
    content: `# Day 6：自动化

今天我们来学习 110Claw 的自动化能力，让 AI 助手在没有你手动触发的情况下自主完成任务。

## 定时任务

使用标准 Cron 表达式创建定时任务：

\`\`\`bash
# 每天早上 8 点发送工作日报
claw schedule "0 8 * * 1-5" "生成昨天的工作总结，发送到我的邮箱"

# 每小时检查服务状态
claw schedule "0 * * * *" "检查生产服务器健康状态，如有异常立即告警"

# 每周一生成周报
claw schedule "0 9 * * 1" "汇总上周的任务完成情况，生成周报"

# 查看所有定时任务
claw schedule list
\`\`\`

## 心跳监测

持续监测系统或服务的状态：

\`\`\`bash
# 监测网站可用性
claw monitor start --url https://mysite.com --interval 5m --alert email

# 监测 API 响应时间
claw monitor start --url https://api.mysite.com/health \
  --interval 1m \
  --threshold "response_time > 500ms"
\`\`\`

## 事件触发

基于特定事件自动执行任务：

\`\`\`bash
# 文件变化时触发
claw watch --path ~/Documents/reports "有新文件时，自动生成摘要并归档"

# Webhook 触发
claw webhook create --name "github-pr" \
  "收到 GitHub PR 时，自动进行代码审查并回复评论"
\`\`\`

## 自动化工作流示例

以下是一个完整的自动化工作流：

\`\`\`yaml
# workflow.yaml
name: daily-digest
triggers:
  - cron: "0 7 * * *"

steps:
  - name: 收集新闻
    action: web-scraper
    params:
      urls: ["https://news.ycombinator.com", "https://v2ex.com"]

  - name: 生成摘要
    action: summarize
    input: "{{steps.收集新闻.output}}"

  - name: 发送邮件
    action: email
    params:
      to: "me@example.com"
      subject: "每日科技摘要"
      body: "{{steps.生成摘要.output}}"
\`\`\`

## 今日练习

1. 创建一个每天早上汇报天气的定时任务
2. 设置一个监控你常用网站的心跳任务
3. 尝试编写一个多步骤的自动化工作流

## 小结

自动化是 AI 助手从"工具"升级为"助理"的关键一步。明天我们将学习最高级的技巧——多 Agent、浏览器控制和设备整合！`,
  },
  {
    day: 7,
    title: '高级技巧',
    desc: '多 Agent、浏览器控制、设备整合，解锁全部潜力。',
    duration: '约 90 分钟',
    content: `# Day 7：高级技巧

恭喜你来到最后一天！今天我们将探索 110Claw 最前沿的高级功能。

## 多 Agent 协作

将复杂任务分解给多个专业 Agent 并行处理：

\`\`\`bash
# 创建多 Agent 工作流
claw agent run --workflow multi-agent.yaml
\`\`\`

\`\`\`yaml
# multi-agent.yaml
name: research-report
agents:
  researcher:
    model: claude-opus
    role: "你是一名资深研究员，专注于收集和验证信息"

  analyst:
    model: claude-sonnet
    role: "你是一名数据分析师，专注于数据解读和洞察"

  writer:
    model: claude-sonnet
    role: "你是一名技术写作专家，专注于清晰的文档输出"

workflow:
  - agent: researcher
    task: "收集关于 {topic} 的最新数据和研究"
  - agent: analyst
    task: "分析 researcher 的发现，提炼关键洞察"
    input: "{{researcher.output}}"
  - agent: writer
    task: "将分析结果写成一份专业报告"
    input: "{{analyst.output}}"
\`\`\`

## 浏览器控制

110Claw 可以控制浏览器，实现真正的网页自动化：

\`\`\`bash
# 安装浏览器控制 Skill
claw skill install browser-control

# 自动化浏览器操作
claw chat "打开 GitHub，搜索 '110claw'，截图第一个结果页面"

# 表单自动填写
claw chat "打开报销系统，将这些费用数据自动填入报销单"
\`\`\`

## 设备整合

### 与智能家居集成

\`\`\`bash
# 安装智能家居 Skill
claw skill install smart-home

# 语音控制
claw chat "下班了，帮我关灯、开空调、拉窗帘"
\`\`\`

### 与手机整合

110Claw 移动端支持快捷指令（iOS Shortcuts / Android Tasker）集成：

\`\`\`bash
# 扫描二维码/文字后自动处理
# 截图后自动分析内容
# 基于位置触发任务
\`\`\`

## 高级提示技巧

### Chain of Thought（思维链）

\`\`\`
请一步一步地思考这个问题：
1. 首先分析问题的本质
2. 列出可能的解决方案
3. 评估每个方案的利弊
4. 给出最终建议和实施步骤
\`\`\`

### 自我反思

\`\`\`
完成任务后，请检查你的回答：
- 是否完整回答了问题？
- 有没有遗漏重要信息？
- 是否有更好的表达方式？
\`\`\`

## 7 天回顾与展望

恭喜你完成了 7 天学习路径！让我们回顾一下学习成果：

| Day | 主题 | 核心技能 |
|-----|------|---------|
| 1 | 初识 110Claw | 安装配置，基础对话 |
| 2 | 深入对话 | 提示词工程，多轮对话 |
| 3 | 文件与代码 | 文件处理，代码生成 |
| 4 | 网络能力 | 搜索，API 调用 |
| 5 | 技能扩展 | Skill 安装，Skill 开发 |
| 6 | 自动化 | 定时任务，工作流 |
| 7 | 高级技巧 | 多 Agent，浏览器控制 |

## 下一步

- 加入 110Claw 社区，分享你的使用经验
- 探索 ClawHub，发现更多有趣的 Skill
- 尝试开发自己的 Skill，为社区做贡献
- 关注 110Claw 官方动态，第一时间获取新功能

**感谢你的坚持，AI 助手的世界等你探索！🎉**`,
  },
]

export function getStepByDay(day) {
  return learningSteps.find(s => s.day === Number(day)) || null
}
