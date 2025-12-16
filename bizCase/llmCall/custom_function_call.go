package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizCase/llmproxy/parser"
	"github.com/morehao/go-action/bizCase/llmproxy/renderer"
	"github.com/morehao/go-action/bizCase/llmproxy/types"
	"github.com/morehao/golib/glog"
	"github.com/openai/openai-go"
)

// executeFunctionCall 根据函数名执行对应的函数
func executeFunctionCall(ctx context.Context, funcName string, args map[string]interface{}) (string, error) {
	glog.Infof(ctx, "[executeFunctionCall] 执行函数: %s, 参数: %v", funcName, args)

	switch funcName {
	case "get_weather":
		// 提取 location 参数
		location, ok := args["location"].(string)
		if !ok {
			return "", fmt.Errorf("缺少必需参数 location 或参数类型错误")
		}
		// 调用天气函数
		result := getWeather(location)
		return result, nil

	// 可以在这里添加更多函数的支持
	// case "other_function":
	//     return handleOtherFunction(args)

	default:
		return "", fmt.Errorf("未知的函数名: %s", funcName)
	}
}

// convertToOpenAIMessages 将 types.Message 转换为 openai.ChatCompletionMessageParamUnion
func convertToOpenAIMessages(messages []types.Message) []openai.ChatCompletionMessageParamUnion {
	result := make([]openai.ChatCompletionMessageParamUnion, len(messages))
	for i, msg := range messages {
		switch msg.Role {
		case "system":
			result[i] = openai.SystemMessage(msg.Content)
		case "user":
			result[i] = openai.UserMessage(msg.Content)
		case "assistant":
			result[i] = openai.AssistantMessage(msg.Content)
		default:
			result[i] = openai.UserMessage(msg.Content)
		}
	}
	return result
}

// CustomFunctionCall 自定义函数调用处理（普通调用）
func CustomFunctionCall(ctx *gin.Context) {
	// 定义可用的工具 - 使用 llmproxy 的标准格式
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
							"description": "城市名称，例如：北京、上海、西安",
						},
					},
					"required": []string{"location"},
				},
			},
		},
	}

	// 用户问题
	userQuestion := "西安今天天气怎么样"

	// 构建请求 - 使用 llmproxy 的 types.ChatRequest
	chatReq := &types.ChatRequest{
		Model: string(LLMModel),
		Messages: []types.Message{
			{Role: "user", Content: userQuestion},
		},
		Tools: tools,
	}

	// 使用 renderer 渲染工具到系统提示词
	render := renderer.NewRenderer("generic") // 使用通用 JSON 格式
	modifiedReq := render.RenderTools(chatReq)

	glog.Info(ctx, "[CustomFunctionCall] 发送第一次请求，询问模型是否需要调用函数")

	// 转换为 OpenAI SDK 格式
	openaiMessages := convertToOpenAIMessages(modifiedReq.Messages)

	// 第一次调用：让模型判断是否需要调用函数
	params := openai.ChatCompletionNewParams{
		Messages: openaiMessages,
		Model:    LLMModel,
		Seed:     openai.Int(0),
	}

	completion, err := llmClient.Chat.Completions.New(ctx, params)
	if err != nil {
		glog.Errorf(ctx, "[CustomFunctionCall] 第一次调用失败: %s", err.Error())
		ctx.JSON(200, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	firstResponse := completion.Choices[0].Message.Content
	glog.Infof(ctx, "[CustomFunctionCall] 模型第一次响应: %s", firstResponse)

	// 使用 parser 解析函数调用
	p := parser.NewParser("json") // 使用 JSON parser
	toolCalls, remainingContent, parseErr := p.Parse(firstResponse)

	if parseErr != nil || len(toolCalls) == 0 {
		// 如果解析失败或没有工具调用，说明模型直接回答了问题
		glog.Infof(ctx, "[CustomFunctionCall] 未检测到函数调用，直接返回响应: %s", firstResponse)
		ctx.JSON(200, Response{
			Success: true,
			Message: firstResponse,
		})
		return
	}

	// 执行第一个工具调用
	funcCall := toolCalls[0]
	glog.Infof(ctx, "[CustomFunctionCall] 检测到函数调用: %s", funcCall.Name)
	funcResult, execErr := executeFunctionCall(ctx, funcCall.Name, funcCall.Arguments)
	if execErr != nil {
		glog.Errorf(ctx, "[CustomFunctionCall] 函数执行失败: %s", execErr.Error())
		ctx.JSON(200, Response{
			Success: false,
			Message: fmt.Sprintf("函数执行失败: %s", execErr.Error()),
		})
		return
	}

	glog.Infof(ctx, "[CustomFunctionCall] 函数执行结果: %s", funcResult)

	// 构建第二次请求的消息
	modifiedReq.Messages = append(modifiedReq.Messages, types.Message{
		Role:    "assistant",
		Content: remainingContent,
	})
	modifiedReq.Messages = append(modifiedReq.Messages, types.Message{
		Role:    "user",
		Content: fmt.Sprintf("函数 %s 的执行结果：%s\n\n请根据这个结果回答用户的问题。", funcCall.Name, funcResult),
	})

	// 第二次调用：让模型基于函数结果生成最终回答
	openaiMessages = convertToOpenAIMessages(modifiedReq.Messages)
	params.Messages = openaiMessages

	glog.Info(ctx, "[CustomFunctionCall] 发送第二次请求，让模型基于函数结果生成回答")

	finalCompletion, finalErr := llmClient.Chat.Completions.New(ctx, params)
	if finalErr != nil {
		glog.Errorf(ctx, "[CustomFunctionCall] 第二次调用失败: %s", finalErr.Error())
		ctx.JSON(200, Response{
			Success: false,
			Message: finalErr.Error(),
		})
		return
	}

	finalResponse := finalCompletion.Choices[0].Message.Content
	glog.Infof(ctx, "[CustomFunctionCall] 最终响应: %s", finalResponse)

	ctx.JSON(200, Response{
		Success: true,
		Message: finalResponse,
	})
}

// CustomStreamFunctionCall 自定义函数调用处理（流式调用）
func CustomStreamFunctionCall(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream; charset=utf-8")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// 定义可用的工具 - 使用 llmproxy 的标准格式
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
							"description": "城市名称，例如：北京、上海、西安",
						},
					},
					"required": []string{"location"},
				},
			},
		},
	}

	// 用户问题
	userQuestion := "西安今天天气怎么样"

	// 构建请求 - 使用 llmproxy 的 types.ChatRequest
	chatReq := &types.ChatRequest{
		Model: string(LLMModel),
		Messages: []types.Message{
			{Role: "user", Content: userQuestion},
		},
		Tools: tools,
	}

	// 使用 renderer 渲染工具到系统提示词
	render := renderer.NewRenderer("generic") // 使用通用 JSON 格式
	modifiedReq := render.RenderTools(chatReq)

	ctx.SSEvent("start", "开始处理请求")
	ctx.Writer.Flush()

	glog.Info(ctx, "[CustomStreamFunctionCall] 发送第一次流式请求")

	// 转换为 OpenAI SDK 格式
	openaiMessages := convertToOpenAIMessages(modifiedReq.Messages)

	// 第一次调用：让模型判断是否需要调用函数（流式）
	params := openai.ChatCompletionNewParams{
		Messages: openaiMessages,
		Model:    LLMModel,
		Seed:     openai.Int(0),
	}

	stream := llmClient.Chat.Completions.NewStreaming(ctx, params)

	// 使用流式解析器
	streamParser := parser.NewStreamParser("json")
	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			content := chunk.Choices[0].Delta.Content
			// 使用流式解析器解析
			_, _, _ = streamParser.Add(content)
		}
	}

	if err := stream.Err(); err != nil {
		glog.Errorf(ctx, "[CustomStreamFunctionCall] 第一次流式调用失败: %s", err.Error())
		ctx.SSEvent("error", err.Error())
		return
	}

	// 刷新解析器，获取完整的工具调用和剩余内容
	toolCalls, remainingContent := streamParser.Flush()

	glog.Infof(ctx, "[CustomStreamFunctionCall] 第一次响应解析完成，工具调用数: %d", len(toolCalls))

	if len(toolCalls) == 0 {
		// 如果没有工具调用，说明模型直接回答了问题
		fullContent := acc.Choices[0].Message.Content
		glog.Infof(ctx, "[CustomStreamFunctionCall] 未检测到函数调用，直接流式返回响应")
		ctx.SSEvent("content", fullContent)
		ctx.Writer.Flush()
		ctx.SSEvent("done", "done")
		ctx.Writer.Flush()
		return
	}

	// 执行第一个工具调用
	funcCall := toolCalls[0]
	glog.Infof(ctx, "[CustomStreamFunctionCall] 检测到函数调用: %s", funcCall.Name)
	ctx.SSEvent("function_call", fmt.Sprintf("正在调用函数: %s", funcCall.Name))
	ctx.Writer.Flush()

	funcResult, execErr := executeFunctionCall(ctx, funcCall.Name, funcCall.Arguments)
	if execErr != nil {
		glog.Errorf(ctx, "[CustomStreamFunctionCall] 函数执行失败: %s", execErr.Error())
		ctx.SSEvent("error", fmt.Sprintf("函数执行失败: %s", execErr.Error()))
		return
	}

	glog.Infof(ctx, "[CustomStreamFunctionCall] 函数执行结果: %s", funcResult)
	ctx.SSEvent("function_result", fmt.Sprintf("函数执行完成: %s", funcCall.Name))
	ctx.Writer.Flush()

	// 构建第二次请求的消息
	modifiedReq.Messages = append(modifiedReq.Messages, types.Message{
		Role:    "assistant",
		Content: remainingContent,
	})
	modifiedReq.Messages = append(modifiedReq.Messages, types.Message{
		Role:    "user",
		Content: fmt.Sprintf("函数 %s 的执行结果：%s\n\n请根据这个结果回答用户的问题。", funcCall.Name, funcResult),
	})

	// 第二次调用：让模型基于函数结果生成最终回答（流式）
	openaiMessages = convertToOpenAIMessages(modifiedReq.Messages)
	params.Messages = openaiMessages

	glog.Info(ctx, "[CustomStreamFunctionCall] 发送第二次流式请求")

	finalStream := llmClient.Chat.Completions.NewStreaming(ctx, params)
	finalAcc := openai.ChatCompletionAccumulator{}

	ctx.SSEvent("answer_start", "开始生成最终回答")
	ctx.Writer.Flush()

	for finalStream.Next() {
		chunk := finalStream.Current()
		finalAcc.AddChunk(chunk)

		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			content := chunk.Choices[0].Delta.Content
			ctx.SSEvent("content", content)
			ctx.Writer.Flush()
		}

		// 判断是否拒绝生成
		if refusal, ok := finalAcc.JustFinishedRefusal(); ok {
			ctx.SSEvent("refusal", refusal)
			ctx.Writer.Flush()
		}
	}

	if err := finalStream.Err(); err != nil {
		glog.Errorf(ctx, "[CustomStreamFunctionCall] 第二次流式调用失败: %s", err.Error())
		ctx.SSEvent("error", err.Error())
		return
	}

	// 发送使用量信息
	if finalAcc.Usage.TotalTokens > 0 {
		usageJSON, _ := json.Marshal(map[string]interface{}{
			"prompt_tokens":     finalAcc.Usage.PromptTokens,
			"completion_tokens": finalAcc.Usage.CompletionTokens,
			"total_tokens":      finalAcc.Usage.TotalTokens,
		})
		ctx.SSEvent("usage", string(usageJSON))
		ctx.Writer.Flush()
	}

	// 发送结束信号
	ctx.SSEvent("done", "done")
	ctx.Writer.Flush()

	glog.Info(ctx, "[CustomStreamFunctionCall] 流式响应完成")
}
