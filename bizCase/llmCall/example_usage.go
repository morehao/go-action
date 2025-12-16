package main

// 本文件展示如何使用 llmproxy 包来实现自定义 function calling
// 这是一个简化的示例，供参考

import (
	"context"
	"fmt"

	"github.com/morehao/go-action/bizCase/llmproxy/parser"
	"github.com/morehao/go-action/bizCase/llmproxy/renderer"
	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// ExampleBasicUsage 展示基本用法
func ExampleBasicUsage() {
	// 1. 定义工具
	tools := []types.Tool{
		{
			Type: "function",
			Function: types.ToolFunction{
				Name:        "get_weather",
				Description: "获取指定位置的天气",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"location": map[string]interface{}{
							"type":        "string",
							"description": "城市名称",
						},
					},
					"required": []string{"location"},
				},
			},
		},
	}

	// 2. 构建请求
	req := &types.ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []types.Message{
			{Role: "user", Content: "北京今天天气怎么样？"},
		},
		Tools: tools,
	}

	// 3. 使用 renderer 渲染工具定义
	render := renderer.NewRenderer("generic") // 或 "qwen", "llama"
	modifiedReq := render.RenderTools(req)

	// 现在 modifiedReq.Messages 的第一条消息包含了工具定义的系统提示词
	fmt.Printf("系统提示词已添加，消息数：%d\n", len(modifiedReq.Messages))

	// 4. 调用模型（这里省略实际调用）
	// response := callLLM(modifiedReq)

	// 5. 解析响应
	modelResponse := `{"name": "get_weather", "arguments": {"location": "北京"}}`

	p := parser.NewParser("json") // 或 "xml"
	toolCalls, remainingContent, err := p.Parse(modelResponse)

	if err == nil && len(toolCalls) > 0 {
		fmt.Printf("检测到工具调用：%s\n", toolCalls[0].Name)
		fmt.Printf("参数：%v\n", toolCalls[0].Arguments)
		fmt.Printf("剩余内容：%s\n", remainingContent)
	}
}

// ExampleStreamUsage 展示流式解析用法
func ExampleStreamUsage() {
	// 创建流式解析器
	streamParser := parser.NewStreamParser("json")

	// 模拟接收流式内容
	chunks := []string{
		`{"name":`,
		` "get_weather"`,
		`, "arguments":`,
		` {"location": "北京"}}`,
	}

	// 逐块添加内容
	for _, chunk := range chunks {
		toolCalls, content, _ := streamParser.Add(chunk)
		if len(toolCalls) > 0 {
			fmt.Printf("检测到工具调用（流式）：%s\n", toolCalls[0].Name)
		}
		if content != "" {
			fmt.Printf("可输出内容：%s\n", content)
		}
	}

	// 刷新缓冲区
	allToolCalls, remaining := streamParser.Flush()
	fmt.Printf("总共检测到 %d 个工具调用\n", len(allToolCalls))
	fmt.Printf("剩余内容：%s\n", remaining)
}

// ExampleDifferentFormats 展示不同格式的支持
func ExampleDifferentFormats() {
	tools := []types.Tool{
		{
			Type: "function",
			Function: types.ToolFunction{
				Name:        "search",
				Description: "搜索信息",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"query": map[string]interface{}{
							"type": "string",
						},
					},
				},
			},
		},
	}

	req := &types.ChatRequest{
		Model:    "qwen-plus",
		Messages: []types.Message{{Role: "user", Content: "搜索天气"}},
		Tools:    tools,
	}

	// 使用不同的 renderer
	formats := []string{"generic", "qwen", "llama"}

	for _, format := range formats {
		render := renderer.NewRenderer(format)
		modifiedReq := render.RenderTools(req)

		// 查看不同格式的系统提示词
		if len(modifiedReq.Messages) > 0 {
			fmt.Printf("\n=== %s 格式 ===\n", format)
			fmt.Printf("系统提示词：\n%s\n", modifiedReq.Messages[0].Content)
		}
	}
}

// ExampleXMLParsing 展示 XML 格式解析
func ExampleXMLParsing() {
	// XML 格式的模型响应
	xmlResponse := `我来帮你查询天气。
<tool_call>
{"name": "get_weather", "arguments": {"location": "上海"}}
</tool_call>
好的，正在查询上海的天气。`

	// 使用 XML parser
	p := parser.NewParser("xml")
	toolCalls, remainingContent, err := p.Parse(xmlResponse)

	if err == nil {
		fmt.Printf("XML 解析成功\n")
		if len(toolCalls) > 0 {
			fmt.Printf("工具调用：%s\n", toolCalls[0].Name)
			fmt.Printf("参数：%v\n", toolCalls[0].Arguments)
		}
		fmt.Printf("剩余文本：%s\n", remainingContent)
	}
}

// ExampleErrorHandling 展示错误处理
func ExampleErrorHandling(ctx context.Context) {
	modelResponse := "这是一个普通回答，没有工具调用"

	p := parser.NewParser("json")
	toolCalls, content, err := p.Parse(modelResponse)

	if err != nil || len(toolCalls) == 0 {
		// 没有工具调用，直接返回内容
		fmt.Printf("直接回答：%s\n", content)
		return
	}

	// 有工具调用，执行工具
	for _, tc := range toolCalls {
		result, err := executeFunctionCall(ctx, tc.Name, tc.Arguments)
		if err != nil {
			fmt.Printf("工具执行失败：%v\n", err)
			continue
		}
		fmt.Printf("工具执行成功：%s\n", result)
	}
}

// 完整的工作流程示例
func ExampleCompleteWorkflow(ctx context.Context, userQuestion string) {
	// 步骤 1: 定义工具
	tools := []types.Tool{
		{
			Type: "function",
			Function: types.ToolFunction{
				Name:        "get_weather",
				Description: "获取天气信息",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"location": map[string]interface{}{
							"type": "string",
						},
					},
					"required": []string{"location"},
				},
			},
		},
	}

	// 步骤 2: 构建并渲染请求
	req := &types.ChatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []types.Message{{Role: "user", Content: userQuestion}},
		Tools:    tools,
	}

	render := renderer.NewRenderer("generic")
	modifiedReq := render.RenderTools(req)

	// 步骤 3: 调用模型（伪代码）
	// firstResponse := callLLM(modifiedReq)
	firstResponse := `{"name": "get_weather", "arguments": {"location": "北京"}}`

	// 步骤 4: 解析响应
	p := parser.NewParser("json")
	toolCalls, remainingContent, err := p.Parse(firstResponse)

	if err != nil || len(toolCalls) == 0 {
		// 没有工具调用，直接返回
		fmt.Printf("最终回答：%s\n", firstResponse)
		return
	}

	// 步骤 5: 执行工具
	funcResult, err := executeFunctionCall(ctx, toolCalls[0].Name, toolCalls[0].Arguments)
	if err != nil {
		fmt.Printf("工具执行失败：%v\n", err)
		return
	}

	// 步骤 6: 构建第二次请求
	modifiedReq.Messages = append(modifiedReq.Messages, types.Message{
		Role:    "assistant",
		Content: remainingContent,
	})
	modifiedReq.Messages = append(modifiedReq.Messages, types.Message{
		Role:    "user",
		Content: fmt.Sprintf("函数执行结果：%s\n\n请根据结果回答问题。", funcResult),
	})

	// 步骤 7: 再次调用模型生成最终回答（伪代码）
	// finalResponse := callLLM(modifiedReq)
	finalResponse := "根据查询结果，北京今天天气晴朗，温度 25°C。"

	fmt.Printf("最终回答：%s\n", finalResponse)
}

// 注意事项：
// 1. 根据实际使用的模型选择合适的 renderer 格式（generic/qwen/llama）
// 2. 根据模型输出选择合适的 parser 格式（json/xml）
// 3. 流式场景使用 StreamParser 可以获得更好的性能
// 4. 始终检查 parser 返回的错误和工具调用数量
// 5. 工具执行失败时要有合适的错误处理
