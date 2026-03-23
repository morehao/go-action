# agentdemo

基于 [Eino](https://github.com/cloudwego/eino) 框架和 [Gin](https://github.com/gin-gonic/gin) 构建的 ReAct Agent HTTP 服务 Demo。

## 项目简介

本 Demo 演示了如何在 Go 项目中结合 **Eino ReAct Agent** 与 **Gin Web 框架**，快速搭建一个具备工具调用能力的智能对话 API 服务。

功能特性：
- 基于 Eino `flow/agent/react` 构建的 ReAct Agent
- 内置两个示例工具：`get_current_time`（获取当前时间）、`calculator`（四则运算）
- 支持 **同步** (`POST /chat`) 和 **流式 SSE** (`POST /chat/stream`) 两种响应模式
- 可通过配置文件切换任意兼容 OpenAI 接口的大语言模型

## 目录结构

```
agentdemo/
├── README.md
├── go.mod
├── main.go               # 程序入口，启动 Gin 服务
├── config/
│   ├── config.go         # 配置结构体与加载逻辑
│   └── config.yaml       # 配置文件（填写 API Key 后即可运行）
├── agent/
│   └── agent.go          # ReAct Agent 初始化
├── handler/
│   └── chat.go           # Gin HTTP 处理器
└── tools/
    └── tools.go          # 工具实现
```

## 快速开始

### 1. 填写配置

编辑 `config/config.yaml`，填入您的 API Key 及模型信息：

```yaml
model:
  api_key: "your-api-key-here"
  model: "deepseek-chat"          # 支持任意 OpenAI 兼容模型
  base_url: "https://api.deepseek.com/v1"

server:
  port: ":8080"
```

> **提示**：`base_url` 可替换为 OpenAI、智谱、百川等任意兼容 OpenAI 接口的地址。

### 2. 安装依赖

```bash
go mod download
```

### 3. 启动服务

```bash
go run main.go
```

服务启动后监听 `http://localhost:8080`。

## API 接口

### 同步对话

```
POST /chat
```

**请求体：**

```json
{
  "messages": [
    { "role": "user", "content": "现在几点了？" }
  ]
}
```

**响应：**

```json
{
  "code": 0,
  "message": "success",
  "data": "现在是 2025-03-23 08:03:00。"
}
```

---

### 流式对话（SSE）

```
POST /chat/stream
```

请求体与 `/chat` 相同。响应为 Server-Sent Events 格式：

```
event: message
data: 现在是

event: message
data:  2025-03-23 08:03:00。

event: done
data: [DONE]
```

## 架构说明

```
HTTP 请求
    │
    ▼
Gin Handler (handler/chat.go)
    │
    ▼
ReAct Agent (agent/agent.go)
 ├─ ChatModel (DeepSeek / OpenAI compatible)
 └─ ToolsNode
      ├─ get_current_time  →  返回当前时间
      └─ calculator        →  四则运算
```

ReAct 循环：
1. 模型根据用户输入决定是否调用工具
2. 若需要工具，ToolsNode 执行并将结果返回给模型
3. 模型综合工具结果生成最终回答
4. 无需工具时直接输出回答
