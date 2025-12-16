# LLM Proxy

一个为不支持原生 function call 的大语言模型提供 function call 能力的 Go 库。完全兼容 OpenAI API 标准。

## 特性

- 🚀 **零依赖**：仅使用 Go 标准库
- 🔌 **OpenAI 兼容**：完全符合 OpenAI Chat Completion API 标准
- 🎯 **智能代理**：自动检测并为不支持原生 function call 的模型提供模拟实现
- 📡 **流式支持**：完整的流式解析和响应
- 🎨 **多格式支持**：支持多种模型的工具调用格式（Generic XML、Qwen、Llama）
- 🔧 **灵活配置**：支持自定义渲染器和解析器

## 架构

```
用户请求 → Proxy → 判断模型能力
                    ├─支持原生 → 直接转发
                    └─不支持 → Renderer → 转换为提示词 → LLM API
                                                        ↓
                                                    响应文本
                                                        ↓
                                                    Parser → 解析工具调用 → OpenAI 格式响应
```

## 安装

```bash
go get github.com/morehao/go-action/bizCase/llmproxy
```

## 快速开始

### 基础用法

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/morehao/go-action/bizCase/llmproxy"
)

func main() {
    // 创建客户端
    client := llmproxy.NewClient("https://api.openai.com/v1", "your-api-key")
    
    // 创建代理（使用通用 XML 格式）
    proxy := llmproxy.NewProxy(client, "generic")
    
    // 定义工具
    weatherTool := llmproxy.Tool{
        Type: "function",
        Function: llmproxy.ToolFunction{
            Name:        "get_weather",
            Description: "获取指定城市的天气信息",
            Parameters: map[string]interface{}{
                "type": "object",
                "properties": map[string]interface{}{
                    "city": map[string]interface{}{
                        "type":        "string",
                        "description": "城市名称",
                    },
                },
                "required": []string{"city"},
            },
        },
    }
    
    // 构建请求
    req := &llmproxy.ChatRequest{
        Model: "gpt-3.5-turbo",
        Messages: []llmproxy.Message{
            {Role: "user", Content: "今天北京的天气怎么样？"},
        },
        Tools: []llmproxy.Tool{weatherTool},
    }
    
    // 调用代理
    // supportsNative=false 表示模型不支持原生 function call
    ctx := context.Background()
    resp, err := proxy.Chat(ctx, req, false)
    if err != nil {
        log.Fatal(err)
    }
    
    // 处理响应
    if len(resp.Choices) > 0 {
        msg := resp.Choices[0].Message
        fmt.Printf("内容: %s\n", msg.Content)
        
        // 检查工具调用
        if len(msg.ToolCalls) > 0 {
            for _, tc := range msg.ToolCalls {
                fmt.Printf("工具: %s\n", tc.Function.Name)
                fmt.Printf("参数: %s\n", tc.Function.Arguments)
            }
        }
    }
}
```

### 流式用法

```go
// 创建流式请求
stream, err := proxy.StreamChat(ctx, req, false)
if err != nil {
    log.Fatal(err)
}

// 处理流式响应
for chunk := range stream {
    if chunk.Error != nil {
        log.Printf("Error: %v", chunk.Error)
        continue
    }
    
    if len(chunk.Choices) > 0 {
        delta := chunk.Choices[0].Delta
        
        // 输出内容
        if delta.Content != "" {
            fmt.Print(delta.Content)
        }
        
        // 处理工具调用
        if len(delta.ToolCalls) > 0 {
            for _, tc := range delta.ToolCalls {
                fmt.Printf("\n[调用工具: %s]\n", tc.Function.Name)
            }
        }
    }
}
```

## 渲染器格式

### Generic（通用格式）

使用 XML 标签格式，适用于大多数模型：

```xml
<tools>
  {"type": "function", "function": {...}}
</tools>

<tool_call>
{"name": "function_name", "arguments": {...}}
</tool_call>
```

### Qwen（Qwen 专用格式）

使用 Qwen 特定的 XML 格式：

```xml
<tools>
  <function>
    <name>function_name</name>
    <description>...</description>
    <parameters>
      <parameter>
        <name>param_name</name>
        <type>string</type>
      </parameter>
    </parameters>
  </function>
</tools>
```

### Llama（Llama 专用格式）

使用 Python 函数签名格式：

```python
def function_name(param1: str, param2: int) -> dict:
    '''Description
    
    Args:
        param1 (str): Parameter description
        param2 (int): Parameter description
    '''
    pass
```

## API 文档

### 核心类型

#### `ChatRequest`

聊天请求，完全兼容 OpenAI API：

```go
type ChatRequest struct {
    Model            string          `json:"model"`
    Messages         []Message       `json:"messages"`
    Tools            []Tool          `json:"tools,omitempty"`
    Stream           bool            `json:"stream"`
    Temperature      *float64        `json:"temperature,omitempty"`
    MaxTokens        *int            `json:"max_tokens,omitempty"`
    // ... 其他 OpenAI 兼容字段
}
```

#### `ChatResponse`

聊天响应：

```go
type ChatResponse struct {
    ID      string   `json:"id"`
    Object  string   `json:"object"`
    Created int64    `json:"created"`
    Model   string   `json:"model"`
    Choices []Choice `json:"choices"`
    Usage   *Usage   `json:"usage,omitempty"`
}
```

#### `Message`

消息：

```go
type Message struct {
    Role       string     `json:"role"`
    Content    string     `json:"content"`
    ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
    ToolCallID string     `json:"tool_call_id,omitempty"`
}
```

### 核心方法

#### `NewClient`

创建 HTTP 客户端：

```go
func NewClient(baseURL, apiKey string) *Client
```

#### `NewProxy`

创建代理实例：

```go
func NewProxy(client *Client, format string) *Proxy
```

参数：
- `client`: HTTP 客户端
- `format`: 渲染器格式（"generic"、"qwen"、"llama"）

#### `Chat`

发送聊天请求：

```go
func (p *Proxy) Chat(ctx context.Context, req *ChatRequest, supportsNative bool) (*ChatResponse, error)
```

参数：
- `ctx`: 上下文
- `req`: 聊天请求
- `supportsNative`: 模型是否支持原生 function call

#### `StreamChat`

发送流式聊天请求：

```go
func (p *Proxy) StreamChat(ctx context.Context, req *ChatRequest, supportsNative bool) (<-chan StreamChunk, error)
```

## 工作原理

### 1. 工具定义转换

对于不支持原生 function call 的模型，Renderer 将工具定义转换为提示词：

```
原始工具定义 → Renderer → 包含工具说明的系统消息 → 发送给模型
```

### 2. 工具调用解析

Parser 从模型输出中提取工具调用：

```
模型输出文本 → Parser → 识别工具调用标签 → 解析 JSON → 转换为 OpenAI 格式
```

### 3. 流式处理

流式解析器使用状态机实现增量解析：

```
状态0: 查找标签 → 状态1: 解析工具调用 → 状态2: 完成
```

## 高级用法

### 自定义渲染器

```go
type CustomRenderer struct{}

func (r *CustomRenderer) RenderTools(req *llmproxy.ChatRequest) *llmproxy.ChatRequest {
    // 自定义工具渲染逻辑
    return req
}
```

### 自定义解析器

```go
type CustomParser struct{}

func (p *CustomParser) Parse(content string) ([]llmproxy.InternalToolCall, string, error) {
    // 自定义解析逻辑
    return toolCalls, remainingContent, nil
}
```

## 示例

完整示例请查看 `examples/` 目录：

- `examples/basic/` - 基础使用示例
- `examples/streaming/` - 流式调用示例

运行示例：

```bash
# 设置环境变量
export OPENAI_API_KEY="your-api-key"
export OPENAI_BASE_URL="https://api.openai.com/v1"

# 运行基础示例
cd examples/basic
go run main.go

# 运行流式示例
cd examples/streaming
go run main.go
```

## 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./parser
go test ./renderer

# 运行测试并显示覆盖率
go test -cover ./...
```

## 支持的模型

理论上支持所有 OpenAI 兼容的 API，包括但不限于：

- ✅ OpenAI (GPT-3.5, GPT-4) - 支持原生 function call
- ✅ Qwen 系列 - 使用模拟 function call
- ✅ Llama 系列 - 使用模拟 function call
- ✅ 其他 OpenAI 兼容的模型

## 故障排除

### 工具调用未被识别

1. 检查模型输出格式是否正确
2. 尝试不同的渲染器格式（generic、qwen、llama）
3. 查看日志确认解析过程

### 流式响应不完整

1. 确保正确处理 channel 中的所有消息
2. 检查是否调用了 `Flush()` 方法（如果使用自定义解析器）

### JSON 解析错误

1. 验证工具定义的 JSON 格式正确
2. 检查模型输出的 JSON 是否有效

## 参考

本项目参考了 [Ollama](https://github.com/ollama/ollama) 的实现：

- 工具调用解析器设计：`tools/tools.go`
- 渲染器实现：`model/renderers/`
- OpenAI 兼容性：`openai/openai.go`

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

## 作者

morehao

