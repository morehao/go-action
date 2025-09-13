# EinoDeer 项目

## 项目简介

EinoDeer 是一个基于 Golang 开发的智能代理系统，它利用大语言模型（LLM）和工具调用能力，实现了一个可以执行各种任务的智能助手。项目主要特点包括：

- 基于 Eino 框架构建的智能代理系统
- 支持浏览器搜索和网页爬虫功能
- 提供代码生成和数据分析能力
- 支持流式响应的聊天接口

## 系统架构

项目主要由以下几个部分组成：

1. **引擎层（engine）**：包含各种角色代理，如研究员、编码员、协调员等
2. **基础设施层（infra）**：提供工具调用、LLM 接口、日志等基础功能
3. **配置层（config）**：管理系统配置信息
4. **提示词（prompts）**：为不同角色代理提供指导提示
5. **API 接口（handler）**：提供 HTTP 接口服务

## 功能特点

- **浏览器搜索**：模拟网络搜索功能，获取相关信息
- **网页爬虫**：模拟网页内容爬取，提取页面信息
- **代码生成**：根据需求生成 Golang 代码示例
- **数据分析**：提供数据分析相关的代码生成能力
- **流式响应**：支持流式输出的聊天接口

## 运行环境要求

- Go 1.20 或更高版本
- 支持 macOS、Linux 和 Windows 系统

## 安装与运行

### 1. 克隆项目

```bash
git clone <项目仓库地址>
cd einodeer
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置设置

编辑 `config/config.yaml` 文件，根据需要修改配置：

```yaml
tools:
  servers:
    browser_search:
      command: "<命令路径>"
      args: ["--arg1", "--arg2"]
    web_crawler:
      command: "<命令路径>"
      args: ["--arg1", "--arg2"]

models:
  default: "<默认模型名称>"
  api_key: "<API密钥>"
  base_url: "<API基础URL>"

settings:
  max_planning_iterations: 3
  max_steps: 10
```

### 4. 运行服务

```bash
go run main.go
```

服务将在 http://localhost:8888 启动

## API 使用

### 聊天接口

```
POST /api/chat/stream
```

请求示例：

```json
{
  "messages": [
    {
      "role": "user",
      "content": "帮我搜索关于人工智能的最新进展"
    }
  ]
}
```

响应为 Server-Sent Events (SSE) 格式的流式数据。

## 开发与扩展

### 添加新工具

1. 在 `infra/tools.go` 中实现新的工具类型
2. 在 `config/config.yaml` 中添加相应配置

### 添加新角色

1. 在 `prompts/` 目录下添加新角色的提示词
2. 在 `engine/` 目录下实现新角色的逻辑

## 许可证

[项目许可证信息]