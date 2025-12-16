package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/morehao/go-action/bizCase/llmproxy"
)

func main() {
	// 从环境变量获取配置
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		apiKey = "your-api-key-here"
	}

	baseURL := os.Getenv("OPENAI_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}

	// 创建客户端和代理
	client := llmproxy.NewClient(baseURL, apiKey)
	proxy := llmproxy.NewProxy(client, "generic")

	// 定义计算工具
	calculatorTool := llmproxy.Tool{
		Type: "function",
		Function: llmproxy.ToolFunction{
			Name:        "calculate",
			Description: "执行数学计算",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"expression": map[string]interface{}{
						"type":        "string",
						"description": "数学表达式，例如：2+2, 10*5",
					},
				},
				"required": []string{"expression"},
			},
		},
	}

	searchTool := llmproxy.Tool{
		Type: "function",
		Function: llmproxy.ToolFunction{
			Name:        "search",
			Description: "在互联网上搜索信息",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"query": map[string]interface{}{
						"type":        "string",
						"description": "搜索关键词",
					},
				},
				"required": []string{"query"},
			},
		},
	}

	// 构建请求
	req := &llmproxy.ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []llmproxy.Message{
			{
				Role:    "user",
				Content: "请帮我计算 123 + 456 的结果，然后搜索一下 Go 语言的最新特性",
			},
		},
		Tools: []llmproxy.Tool{calculatorTool, searchTool},
	}

	// 调用流式代理
	ctx := context.Background()
	stream, err := proxy.StreamChat(ctx, req, false)
	if err != nil {
		log.Fatalf("Error calling streaming proxy: %v", err)
	}

	// 处理流式响应
	fmt.Println("=== 流式响应 ===")
	fmt.Print("内容: ")

	var fullContent string
	var toolCalls []llmproxy.ToolCall
	var finishReason string

	for chunk := range stream {
		if chunk.Error != nil {
			log.Printf("Stream error: %v", chunk.Error)
			continue
		}

		if len(chunk.Choices) > 0 {
			choice := chunk.Choices[0]

			// 输出内容
			if choice.Delta.Content != "" {
				fmt.Print(choice.Delta.Content)
				fullContent += choice.Delta.Content
			}

			// 收集工具调用
			if len(choice.Delta.ToolCalls) > 0 {
				toolCalls = append(toolCalls, choice.Delta.ToolCalls...)
			}

			// 记录结束原因
			if choice.FinishReason != "" {
				finishReason = choice.FinishReason
			}
		}
	}

	fmt.Println("\n")

	// 显示完整结果
	fmt.Println("\n=== 完整内容 ===")
	fmt.Printf("内容: %s\n", fullContent)
	fmt.Printf("Finish Reason: %s\n", finishReason)

	// 显示工具调用
	if len(toolCalls) > 0 {
		fmt.Println("\n=== 工具调用 ===")
		for i, tc := range toolCalls {
			fmt.Printf("\n调用 %d:\n", i+1)
			fmt.Printf("  ID: %s\n", tc.ID)
			fmt.Printf("  类型: %s\n", tc.Type)
			fmt.Printf("  函数: %s\n", tc.Function.Name)
			fmt.Printf("  参数: %s\n", tc.Function.Arguments)

			// 解析参数
			var args map[string]interface{}
			if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err == nil {
				fmt.Println("  解析后的参数:")
				for k, v := range args {
					fmt.Printf("    %s: %v\n", k, v)
				}
			}

			// 模拟执行工具
			fmt.Println("\n  执行结果:")
			result := executeTool(tc.Function.Name, tc.Function.Arguments)
			fmt.Printf("  %s\n", result)
		}
	}

	// 如果有工具调用，可以继续对话
	if len(toolCalls) > 0 {
		fmt.Println("\n=== 后续操作 ===")
		fmt.Println("提示：您可以将工具执行结果添加到消息中，继续对话")
		fmt.Println("示例代码：")
		fmt.Println(`
for _, tc := range toolCalls {
    result := executeTool(tc.Function.Name, tc.Function.Arguments)
    req.Messages = append(req.Messages, llmproxy.Message{
        Role:       "tool",
        Content:    result,
        ToolCallID: tc.ID,
    })
}
// 继续调用 proxy.StreamChat(ctx, req, false)
`)
	}
}

// executeTool 模拟执行工具
func executeTool(funcName, argsJSON string) string {
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
		return fmt.Sprintf("Error parsing arguments: %v", err)
	}

	switch funcName {
	case "calculate":
		expr, _ := args["expression"].(string)
		// 简单模拟计算（实际应该使用表达式求值器）
		return fmt.Sprintf("计算 %s 的结果为: 579", expr)

	case "search":
		query, _ := args["query"].(string)
		return fmt.Sprintf("搜索 '%s' 的结果：\n1. Go 1.21 引入了泛型优化\n2. Go 1.22 改进了性能\n3. Go 1.23 即将发布", query)

	default:
		return fmt.Sprintf("未知的工具: %s", funcName)
	}
}

