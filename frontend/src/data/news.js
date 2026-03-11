export const newsItems = [
  {
    id: 'news-1',
    title: 'OpenClaw 3.0 正式发布：全新架构，性能提升 300%',
    summary: 'OpenClaw 3.0 带来了全新的插件系统架构，大幅提升了执行效率和稳定性，同时引入了全新的多 Agent 协作机制。',
    content: `OpenClaw 3.0 是迄今为止最重大的版本更新。新版本完全重写了核心调度引擎，引入了基于事件驱动的异步架构，使得 Skill 插件的并发执行效率提升了 300%。

## 主要变更

- **全新插件系统**：Skill 注册表重构，支持热加载和版本锁定
- **多 Agent 协作**：内置 Agent 间通信协议，可以构建复杂的自动化工作流
- **安全沙箱**：每个 Skill 在独立的沙箱环境中执行，防止恶意代码影响主进程
- **内存优化**：空闲时内存占用降低 60%，适合在低配置机器上运行

## 升级指南

从 2.x 升级到 3.0 需要重新安装所有 Skill 插件。大多数常用 Skill 已发布 3.0 兼容版本。详细升级文档请参考官方迁移指南。`,
    date: '2026-03-10',
    tags: ['版本发布', '性能优化'],
    source: '官方公告',
    author: '110Claw 官方',
  },
  {
    id: 'news-2',
    title: '安全警告：发现恶意 Skill "crypto-miner"，请立即卸载',
    summary: '社区发现恶意插件利用用户设备进行加密货币挖矿，已从 ClawHub 移除，请用户立即检查并卸载。',
    content: `安全团队在例行审查中发现，ClawHub 上名为 "crypto-miner" 的 Skill 插件包含恶意代码，会在后台利用用户 CPU 进行加密货币挖矿。

## 影响范围

该插件于 2026-02-20 上架，累计安装量约 2,300 次。

## 建议操作

1. 执行 \`claw skill list\` 查看已安装插件列表
2. 如发现 crypto-miner，立即执行 \`claw skill remove crypto-miner\`
3. 检查系统 CPU 使用率是否异常

## 后续措施

ClawHub 将于近期上线所有插件的自动安全扫描机制，从源头防范恶意代码入侵。`,
    date: '2026-03-08',
    tags: ['安全', 'ClawHub'],
    source: '安全公告',
    author: '110Claw 安全团队',
  },
  {
    id: 'news-3',
    title: 'ClawHub 注册 Skill 数量突破 1000 大关',
    summary: 'OpenClaw 生态持续繁荣，社区贡献的插件数量快速增长，覆盖开发、效率、娱乐等多个领域。',
    content: `截至 2026 年 3 月初，ClawHub 上注册的 Skill 插件数量已突破 1000 个里程碑。这标志着 OpenClaw 生态系统进入了一个崭新阶段。

## 生态数据

- 总 Skill 数量：1,037 个
- 经 110Claw 认证的 Skill：286 个
- 本月新增：142 个
- 最受欢迎类别：编程助手（278个）、数据分析（156个）、自动化工具（134个）

## 质量保障

为维护生态健康，110Claw 团队将加大对优质 Skill 的认证力度，并推出社区评分机制，帮助用户找到真正好用的插件。`,
    date: '2026-03-05',
    tags: ['生态', 'ClawHub'],
    source: '社区动态',
    author: '110Claw 社区',
  },
  {
    id: 'news-4',
    title: '教程：用 110Claw + GitHub Actions 实现代码自动审查',
    summary: '本教程介绍如何将 110Claw 与 GitHub Actions 集成，实现每次 PR 自动进行 AI 代码审查。',
    content: `自动化代码审查可以帮助团队在合并代码前发现潜在问题。本教程将带你一步步配置 110Claw 与 GitHub Actions 的集成。

## 前置条件

- 已安装 110Claw 并配置 API Key
- GitHub 仓库有 Actions 权限
- 安装 \`code-reviewer\` Skill

## 配置步骤

### 1. 创建 GitHub Actions Workflow

\`\`\`yaml
name: AI Code Review
on:
  pull_request:
    types: [opened, synchronize]

jobs:
  review:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Run 110Claw Review
        env:
          CLAW_API_KEY: \${{ secrets.CLAW_API_KEY }}
        run: |
          claw skill run code-reviewer --diff=$(git diff origin/main)
\`\`\`

### 2. 配置 Secrets

在仓库设置中添加 \`CLAW_API_KEY\`。

代码审查结果将自动以 PR 评论形式呈现。`,
    date: '2026-03-03',
    tags: ['教程', '自动化'],
    source: '技术博客',
    author: '社区贡献者 @devhawk',
  },
  {
    id: 'news-5',
    title: '110Claw 移动端体验大幅提升：新版 iOS/Android 客户端发布',
    summary: '全新移动端客户端支持 Skill 插件运行、语音输入、历史对话同步等功能，移动办公更便利。',
    content: `110Claw 全新移动端客户端（iOS 2.1 / Android 2.1）正式上线，带来了大量用户期待已久的功能。

## 新功能亮点

- **Skill 插件支持**：移动端现在可以运行超过 200 个兼容的 Skill 插件
- **语音输入**：支持中文语音转文字，识别准确率达 98%
- **多端同步**：对话历史、Skill 配置自动跨设备同步
- **离线模式**：常用对话模板可缓存，弱网环境下也能正常使用

## 下载地址

- iOS: App Store 搜索 "110Claw"
- Android: Google Play 或官网直接下载 APK`,
    date: '2026-02-28',
    tags: ['版本发布', '移动端'],
    source: '产品公告',
    author: '110Claw 产品团队',
  },
  {
    id: 'news-6',
    title: '深度解析：110Claw 的多 Agent 架构设计',
    summary: '本文深入解析 110Claw 3.0 中多 Agent 协作系统的技术实现，包括消息路由、状态共享与错误隔离机制。',
    content: `多 Agent 系统是 110Claw 3.0 最令人兴奋的新特性之一。本文将从技术角度深入解析其实现原理。

## 架构概述

多 Agent 系统基于发布/订阅模型构建，每个 Agent 都是独立的执行单元，通过消息总线进行通信。

## 核心概念

**Orchestrator Agent**：负责任务分解和工作流调度
**Worker Agent**：执行具体任务的专用 Agent
**Memory Agent**：负责跨 Agent 的状态共享和持久化

## 实际案例

以自动报告生成为例，Orchestrator 将任务分解为数据收集、分析、撰写三个子任务，分别派发给不同的 Worker Agent 并行执行，最后合并结果。整个过程比单 Agent 方案快 4 倍。`,
    date: '2026-02-25',
    tags: ['技术', '架构'],
    source: '技术深度',
    author: '110Claw 核心团队',
  },
  {
    id: 'news-7',
    title: '社区月报：2 月份优质 Skill 推荐',
    summary: '本月社区评选出 10 款最受好评的 Skill 插件，涵盖数据分析、内容创作、系统运维等多个场景。',
    content: `每月一次的社区 Skill 推荐来啦！本月我们从 142 个新增 Skill 中精选出 10 款值得推荐的优质插件。

## 本月推荐

1. **data-viz** - 数据可视化，自动生成图表（⭐ 4.9）
2. **doc-writer** - 代码文档自动生成（⭐ 4.8）
3. **log-analyzer** - 智能日志分析与异常检测（⭐ 4.8）
4. **email-compose** - AI 邮件撰写助手（⭐ 4.7）
5. **stock-watch** - 股票行情监控与分析（⭐ 4.7）

完整榜单及安装方法请访问 ClawHub 月度推荐页面。`,
    date: '2026-02-20',
    tags: ['社区', '推荐'],
    source: '社区月报',
    author: '110Claw 社区编辑',
  },
  {
    id: 'news-8',
    title: '隐私更新：110Claw 全面支持本地模型，数据不出设备',
    summary: '110Claw 现已支持接入本地部署的开源大模型（Ollama/LM Studio），对话数据完全本地化处理。',
    content: `隐私保护是 110Claw 一直以来的核心承诺。全新的本地模型支持功能让用户可以完全掌控自己的数据。

## 支持的本地模型框架

- Ollama（推荐，支持 Llama 3、Qwen2.5 等主流模型）
- LM Studio
- 任何兼容 OpenAI API 格式的本地推理服务

## 配置方法

在 110Claw 设置中选择"本地模型"，填写本地服务地址（通常为 http://localhost:11434），即可完成配置。

## 隐私保障

使用本地模型时，所有对话内容仅在本地处理，不会发送至任何服务器。Skill 插件也可配置为本地优先模式。`,
    date: '2026-02-15',
    tags: ['隐私', '本地模型'],
    source: '产品公告',
    author: '110Claw 产品团队',
  },
  {
    id: 'news-9',
    title: '110Claw 与国内主流 IDE 实现深度集成',
    summary: 'JetBrains 系列、VS Code、Cursor 等主流 IDE 现已支持 110Claw 插件，代码编写体验升级。',
    content: `110Claw IDE 集成插件正式发布，覆盖国内开发者最常用的主流 IDE。

## 支持的 IDE

- **VS Code**：插件市场搜索 "110Claw" 即可安装
- **JetBrains 全家桶**：IntelliJ IDEA、PyCharm、WebStorm 等均已支持
- **Cursor**：原生集成，无需额外配置

## 功能特性

- 代码补全增强（比传统补全更理解上下文）
- 选中代码一键解释/重构/生成注释
- 内联 Chat，无需切换窗口即可与 AI 对话
- 与项目已有 Skill 配置自动同步`,
    date: '2026-02-10',
    tags: ['工具集成', '开发效率'],
    source: '产品公告',
    author: '110Claw 产品团队',
  },
  {
    id: 'news-10',
    title: '开源计划：110Claw 核心协议层将于 Q2 对外开源',
    summary: '110Claw 宣布将于 2026 年 Q2 开源其核心通信协议层，欢迎社区参与共建下一代 AI 助手生态。',
    content: `110Claw 今日宣布将于 2026 年第二季度开源其核心通信协议层（Claw Protocol），这将是 110Claw 生态建设的重大里程碑。

## 开源内容

- **Claw Protocol**：Agent 间通信的标准化协议规范
- **Skill SDK**：Skill 插件开发工具包，包含 Go/Python/TypeScript 三个语言版本
- **CLI 工具链**：完整的命令行工具，包括 Skill 打包、测试、发布全流程

## 社区共建计划

- GitHub 组织：github.com/110claw（即将开放）
- 贡献者激励计划：积分制，优质贡献可兑换会员权益
- RFC 流程：重大特性变更通过社区 RFC 讨论决策

我们相信，开放的协议标准将推动整个 AI 助手生态的繁荣发展。`,
    date: '2026-02-05',
    tags: ['开源', '生态'],
    source: '官方公告',
    author: '110Claw 官方',
  },
]

export const newsTags = ['全部', '版本发布', '安全', '教程', '生态', '技术', '社区', '隐私', '工具集成', '开源', '自动化', '移动端']

export function getNewsById(id) {
  return newsItems.find(n => n.id === id) || null
}
