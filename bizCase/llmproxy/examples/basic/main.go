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

	// 定义天气工具
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
						"description": "城市名称，例如：北京、上海",
					},
					"unit": map[string]interface{}{
						"type":        "string",
						"description": "温度单位",
						"enum":        []string{"celsius", "fahrenheit"},
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
			{
				Role:    "user",
				Content: "今天北京的天气怎么样？",
			},
		},
		Tools: []llmproxy.Tool{weatherTool},
	}

	// 调用代理
	// supportsNative=false 表示模型不支持原生 function call，需要模拟
	ctx := context.Background()
	resp, err := proxy.Chat(ctx, req, false)
	if err != nil {
		log.Fatalf("Error calling proxy: %v", err)
	}

	// 处理响应
	fmt.Println("=== 响应 ===")
	fmt.Printf("Model: %s\n", resp.Model)
	fmt.Printf("ID: %s\n", resp.ID)

	if len(resp.Choices) > 0 {
		choice := resp.Choices[0]
		fmt.Printf("\n内容: %s\n", choice.Message.Content)

		// 检查是否有工具调用
		if len(choice.Message.ToolCalls) > 0 {
			fmt.Println("\n=== 工具调用 ===")
			for i, tc := range choice.Message.ToolCalls {
				fmt.Printf("\n调用 %d:\n", i+1)
				fmt.Printf("  ID: %s\n", tc.ID)
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

				// 模拟执行工具调用
				fmt.Println("\n  执行结果:")
				result := executeWeatherTool(tc.Function.Name, tc.Function.Arguments)
				fmt.Printf("  %s\n", result)
			}
		}

		fmt.Printf("\nFinish Reason: %s\n", choice.FinishReason)
	}

	// 显示 token 使用情况
	if resp.Usage != nil {
		fmt.Println("\n=== Token 使用情况 ===")
		fmt.Printf("Prompt Tokens: %d\n", resp.Usage.PromptTokens)
		fmt.Printf("Completion Tokens: %d\n", resp.Usage.CompletionTokens)
		fmt.Printf("Total Tokens: %d\n", resp.Usage.TotalTokens)
	}
}

// executeWeatherTool 模拟执行天气工具
func executeWeatherTool(funcName, argsJSON string) string {
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
		return fmt.Sprintf("Error parsing arguments: %v", err)
	}

	city, _ := args["city"].(string)
	unit := "celsius"
	if u, ok := args["unit"].(string); ok {
		unit = u
	}

	// 模拟天气数据
	temp := "22"
	if unit == "fahrenheit" {
		temp = "72"
	}

	return fmt.Sprintf("城市 %s 的天气：晴朗，温度 %s°%s，湿度 65%%",
		city, temp, map[string]string{"celsius": "C", "fahrenheit": "F"}[unit])
}

