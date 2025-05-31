package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const (
	ModuleName = "LLM_call"

	LLMHost   = "https://api.deepseek.com/v1"
	LLMModel  = "deepseek-chat"
	LLMAPIKey = "xxxx"
)

var llmClient = openai.NewClient(
	option.WithBaseURL(LLMHost),
	option.WithAPIKey(LLMAPIKey))

func NewChat(ctx *gin.Context) {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	chatCompletion, err := llmClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("你好，我是测试"),
		},
		Model: LLMModel,
	})
	if err != nil {
		ctx.JSON(200, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, Response{
		Success: true,
		Message: chatCompletion.Choices[0].Message.Content,
	})

}
