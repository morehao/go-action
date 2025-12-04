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
- **Prompt 工程**：将 OpenAI 的 tools 参数转换为结构化的文本描述
- **智能解析**：自动解析模型输出的函数调用 JSON 格式
- **自动执行**：根据函数名自动执行对应的函数
- **二次调用**：将函数执行结果返回给模型，生成最终答案
- **OpenAI 协议兼容**：响应格式符合 OpenAI 协议规范

### 扩展新函数：
在 `executeFunctionCall` 函数中添加新的 case 分支即可支持更多函数。