package main

import (
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
)

// 流式调用示例
func StreamChat(ctx *gin.Context) {
	// 设置 SSE 响应头
	ctx.Header("Content-Type", "text/event-stream; charset=utf-8")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	question := "世界上海拔排名前十的山峰"
	sysPrompt := "请简单给出排名即可"
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(sysPrompt),
		openai.UserMessage(question),
	}

	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Seed:     openai.Int(0),
		Model:    LLMModel,
	}

	// 创建流式请求
	stream := llmClient.Chat.Completions.NewStreaming(ctx, params)
	acc := openai.ChatCompletionAccumulator{}
	ctx.SSEvent("start", "start")
	ctx.Writer.Flush()
	for stream.Next() {
		chunk := stream.Current()

		acc.AddChunk(chunk)

		// 若本轮 chunk 含有 Content 增量，则输出
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			content := chunk.Choices[0].Delta.Content
			ctx.SSEvent("content", content)
			ctx.Writer.Flush()
		}

		// 判断是否拒绝生成（如违反政策）
		if refusal, ok := acc.JustFinishedRefusal(); ok {
			ctx.SSEvent("refusal", refusal)
			ctx.Writer.Flush()
		}
	}

	if err := stream.Err(); err != nil {
		ctx.SSEvent("error", err.Error())
		return
	}
	if acc.Usage.TotalTokens > 0 {
		ctx.SSEvent("usage", acc.Usage.TotalTokens)
	}

	// 发送结束信号
	ctx.SSEvent("done", "done")
	ctx.Writer.Flush()
}
