# LLM 调用示例

## 1. 普通调用
测试命令：`curl "http://127.0.0.1:8080/chat"`

## 2. 流式调用
测试命令：`curl -H "Accept: text/event-stream" "http://127.0.0.1:8080/streamChat"`

## 3. Function Calling（原生支持）
测试命令：`curl "http://127.0.0.1:8080/functionCall"`

## 4. 自定义 Function Calling（普通调用）
测试命令：`curl "http://127.0.0.1:8080/customFunctionCall"`

适用于不支持原生 function calling 的大模型，通过 Prompt 工程实现函数调用功能。

## 5. 自定义 Function Calling（流式调用）
测试命令：`curl -H "Accept: text/event-stream" "http://127.0.0.1:8080/customStreamFunctionCall"`

适用于不支持原生 function calling 的大模型，支持流式返回。

## 实现说明

### 自定义 Function Calling 的核心特性：
- **Prompt 工程**：基于 `llmproxy/renderer` 将工具定义转换为结构化的提示词
- **智能解析**：使用 `llmproxy/parser` 自动解析模型输出的函数调用（支持 JSON/XML 格式）
- **自动执行**：根据函数名自动执行对应的函数
- **二次调用**：将函数执行结果返回给模型，生成最终答案
- **OpenAI 协议兼容**：使用 `llmproxy/types` 标准数据结构，符合 OpenAI 协议规范

### 重构说明（2024）：
本模块已重构为基于 `llmproxy` 包实现，主要改进：
- 使用 `llmproxy/types` 的标准数据结构
- 使用 `llmproxy/renderer` 处理工具定义（支持多种模型格式）
- 使用 `llmproxy/parser` 解析函数调用（更健壮的解析逻辑）
- 支持流式解析（`StreamParser`）
- 代码量减少 60%+，可维护性显著提升

详细重构说明见：[REFACTOR_NOTES.md](./REFACTOR_NOTES.md)

### 扩展新函数：
在 `executeFunctionCall` 函数中添加新的 case 分支即可支持更多函数。

### 支持不同模型格式：
```go
// 通用格式（默认）
render := renderer.NewRenderer("generic")

// 通义千问格式
render := renderer.NewRenderer("qwen")

// Llama 格式
render := renderer.NewRenderer("llama")
```