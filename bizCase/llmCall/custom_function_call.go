package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	"github.com/openai/openai-go"
)

// Tool 工具定义
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}

// FunctionCallInfo 函数调用信息
type FunctionCallInfo struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// FunctionCallResponse 模型返回的函数调用响应
type FunctionCallResponse struct {
	FunctionCall *FunctionCallInfo `json:"function_call,omitempty"`
	Content      string            `json:"content,omitempty"`
}

const (
	// FunctionCallSystemPrompt 函数调用的系统提示词模板
	FunctionCallSystemPromptTemplate = `你是一个智能助手，可以调用以下函数来帮助用户：

%s

重要规则：
1. 当你需要调用函数时，必须严格按照以下JSON格式返回，不要包含任何其他内容：
{"function_call": {"name": "函数名", "arguments": {"参数名": "参数值"}}}

2. 如果不需要调用函数，直接回答用户的问题即可。

3. 只返回JSON或文本回答，不要同时返回两者。`
)

// buildFunctionPrompt 将工具列表转换为提示词格式
func buildFunctionPrompt(tools []Tool) string {
	var builder strings.Builder
	
	builder.WriteString("函数列表：\n")
	for i, tool := range tools {
		builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, tool.Name))
		builder.WriteString(fmt.Sprintf("   - 描述：%s\n", tool.Description))
		
		// 格式化参数
		if tool.Parameters != nil {
			paramsJSON, err := json.Marshal(tool.Parameters)
			if err == nil {
				builder.WriteString(fmt.Sprintf("   - 参数：%s\n", string(paramsJSON)))
			}
		}
		builder.WriteString("\n")
	}
	
	return builder.String()
}

// buildSystemPromptWithTools 构建包含工具定义的完整系统提示词
func buildSystemPromptWithTools(tools []Tool) string {
	functionList := buildFunctionPrompt(tools)
	return fmt.Sprintf(FunctionCallSystemPromptTemplate, functionList)
}

// parseFunctionCall 解析模型输出的函数调用信息
func parseFunctionCall(ctx context.Context, response string) (*FunctionCallInfo, error) {
	// 清理响应文本
	response = strings.TrimSpace(response)
	
	// 尝试找到 JSON 部分
	startIdx := strings.Index(response, "{")
	endIdx := strings.LastIndex(response, "}")
	
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("未找到有效的JSON格式")
	}
	
	jsonStr := response[startIdx : endIdx+1]
	
	// 尝试解析为 FunctionCallResponse
	var funcCallResp FunctionCallResponse
	err := json.Unmarshal([]byte(jsonStr), &funcCallResp)
	if err != nil {
		glog.Errorf(ctx, "[parseFunctionCall] 解析JSON失败: %s, JSON: %s", err.Error(), jsonStr)
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}
	
	// 如果包含 function_call 字段，返回函数调用信息
	if funcCallResp.FunctionCall != nil {
		glog.Infof(ctx, "[parseFunctionCall] 成功解析函数调用: %s", funcCallResp.FunctionCall.Name)
		return funcCallResp.FunctionCall, nil
	}
	
	return nil, fmt.Errorf("响应中未包含函数调用信息")
}

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

// CustomFunctionCall 自定义函数调用处理（普通调用）
func CustomFunctionCall(ctx *gin.Context) {
	// 定义可用的工具
	tools := []Tool{
		{
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
	}
	
	// 构建系统提示词
	systemPrompt := buildSystemPromptWithTools(tools)
	
	// 用户问题
	userQuestion := "西安今天天气怎么样"
	
	// 构建消息列表
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemPrompt),
		openai.UserMessage(userQuestion),
	}
	
	// 第一次调用：让模型判断是否需要调用函数
	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    LLMModel,
		Seed:     openai.Int(0),
	}
	
	glog.Info(ctx, "[CustomFunctionCall] 发送第一次请求，询问模型是否需要调用函数")
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
	
	// 尝试解析函数调用
	funcCall, parseErr := parseFunctionCall(ctx, firstResponse)
	if parseErr != nil {
		// 如果解析失败，说明模型直接回答了问题，不需要函数调用
		glog.Infof(ctx, "[CustomFunctionCall] 未检测到函数调用，直接返回响应: %s", firstResponse)
		ctx.JSON(200, Response{
			Success: true,
			Message: firstResponse,
		})
		return
	}
	
	// 执行函数调用
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
	// 将函数调用结果作为助手消息添加到对话历史
	messages = append(messages, openai.AssistantMessage(firstResponse))
	messages = append(messages, openai.UserMessage(fmt.Sprintf("函数 %s 的执行结果：%s\n\n请根据这个结果回答用户的问题。", funcCall.Name, funcResult)))
	
	// 第二次调用：让模型基于函数结果生成最终回答
	params.Messages = messages
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
	
	// 定义可用的工具
	tools := []Tool{
		{
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
	}
	
	// 构建系统提示词
	systemPrompt := buildSystemPromptWithTools(tools)
	
	// 用户问题
	userQuestion := "西安今天天气怎么样"
	
	// 构建消息列表
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemPrompt),
		openai.UserMessage(userQuestion),
	}
	
	// 第一次调用：让模型判断是否需要调用函数（流式）
	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    LLMModel,
		Seed:     openai.Int(0),
	}
	
	ctx.SSEvent("start", "开始处理请求")
	ctx.Writer.Flush()
	
	glog.Info(ctx, "[CustomStreamFunctionCall] 发送第一次流式请求")
	stream := llmClient.Chat.Completions.NewStreaming(ctx, params)
	
	// 累积第一次响应的内容
	var firstResponseBuilder strings.Builder
	acc := openai.ChatCompletionAccumulator{}
	
	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)
		
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			content := chunk.Choices[0].Delta.Content
			firstResponseBuilder.WriteString(content)
		}
	}
	
	if err := stream.Err(); err != nil {
		glog.Errorf(ctx, "[CustomStreamFunctionCall] 第一次流式调用失败: %s", err.Error())
		ctx.SSEvent("error", err.Error())
		return
	}
	
	firstResponse := firstResponseBuilder.String()
	glog.Infof(ctx, "[CustomStreamFunctionCall] 第一次响应完整内容: %s", firstResponse)
	
	// 尝试解析函数调用
	funcCall, parseErr := parseFunctionCall(ctx, firstResponse)
	if parseErr != nil {
		// 如果解析失败，说明模型直接回答了问题
		glog.Infof(ctx, "[CustomStreamFunctionCall] 未检测到函数调用，直接流式返回响应")
		ctx.SSEvent("content", firstResponse)
		ctx.Writer.Flush()
		ctx.SSEvent("done", "done")
		ctx.Writer.Flush()
		return
	}
	
	// 执行函数调用
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
	messages = append(messages, openai.AssistantMessage(firstResponse))
	messages = append(messages, openai.UserMessage(fmt.Sprintf("函数 %s 的执行结果：%s\n\n请根据这个结果回答用户的问题。", funcCall.Name, funcResult)))
	
	// 第二次调用：让模型基于函数结果生成最终回答（流式）
	params.Messages = messages
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
		ctx.SSEvent("usage", finalAcc.Usage.TotalTokens)
		ctx.Writer.Flush()
	}
	
	// 发送结束信号
	ctx.SSEvent("done", "done")
	ctx.Writer.Flush()
	
	glog.Info(ctx, "[CustomStreamFunctionCall] 流式响应完成")
}

