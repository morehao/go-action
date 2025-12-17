# JSON Parser 使用示例

## 功能特性

改进后的 JSON Parser 支持以下格式：

1. **纯 JSON 格式**
2. **Markdown 代码块格式**（带或不带 `json` 标签）
3. **文本 + JSON 混合格式**
4. **Arguments 为字符串的情况**（自动解析为对象）
5. **嵌套 JSON 对象**

## 使用示例

### 示例 1: Markdown 代码块 + 字符串 Arguments

```go
package main

import (
    "fmt"
    "github.com/morehao/go-action/bizCase/llmproxy/parser"
)

func main() {
    // LLM 返回的响应（包含 markdown 代码块和字符串形式的 arguments）
    input := `即将使用获取指定位置的天气
` + "```json\n{\"name\": \"get_weather\", \"arguments\": \"{\\\"location\\\": \\\"西安\\\"}\"}\n```"

    p := parser.NewJSONParser()
    toolCalls, remainingContent, err := p.Parse(input)
    
    if err != nil {
        panic(err)
    }
    
    if len(toolCalls) > 0 {
        fmt.Printf("工具名称: %s\n", toolCalls[0].Name)
        fmt.Printf("工具参数: %+v\n", toolCalls[0].Arguments)
        fmt.Printf("剩余文本: %s\n", remainingContent)
    }
}

// 输出:
// 工具名称: get_weather
// 工具参数: map[location:西安]
// 剩余文本: 即将使用获取指定位置的天气
```

### 示例 2: 纯 JSON 格式（对象 Arguments）

```go
input := `{"name": "get_weather", "arguments": {"location": "北京"}}`

p := parser.NewJSONParser()
toolCalls, remainingContent, err := p.Parse(input)

// 工具名称: get_weather
// 工具参数: map[location:北京]
```

### 示例 3: 文本 + JSON 混合格式

```go
input := `让我查询一下天气 {"name": "get_weather", "arguments": {"location": "上海"}} 好的`

p := parser.NewJSONParser()
toolCalls, remainingContent, err := p.Parse(input)

// 工具名称: get_weather
// 工具参数: map[location:上海]
// 剩余文本: 让我查询一下天气  好的
```

### 示例 4: 嵌套 JSON 对象

```go
input := `{"name": "complex_tool", "arguments": {"nested": {"key": "value"}, "array": [1, 2, 3]}}`

p := parser.NewJSONParser()
toolCalls, remainingContent, err := p.Parse(input)

// 工具名称: complex_tool
// 工具参数: map[array:[1 2 3] nested:map[key:value]]
```

### 示例 5: 流式解析

```go
streamParser := parser.NewStreamJSONParser()

chunks := []string{
    "即将使用获取指定位置的天气\n```json\n",
    "{\"name\": \"get_weather\", ",
    "\"arguments\": \"{\\\"location\\\": \\\"西安\\\"}\"}",
    "\n```",
}

for _, chunk := range chunks {
    toolCalls, content, _ := streamParser.Add(chunk)
    if len(toolCalls) > 0 {
        fmt.Printf("检测到工具调用: %s\n", toolCalls[0].Name)
    }
    if content != "" {
        fmt.Printf("内容: %s\n", content)
    }
}

// 刷新缓冲区
finalToolCalls, finalContent := streamParser.Flush()
fmt.Printf("剩余内容: %s\n", finalContent)
```

## 关键改进

1. **Markdown 代码块支持**: 自动识别 ` ```json ... ``` ` 和 ` ``` ... ``` ` 格式
2. **字符串 Arguments 解析**: 如果 `arguments` 是 JSON 字符串，自动解析为对象
3. **括号匹配算法**: 使用括号计数而非正则，支持任意深度的嵌套 JSON
4. **更好的错误处理**: 解析失败时返回原始内容，不会丢失数据
5. **流式解析去重**: 避免在流式场景中重复检测相同的工具调用
