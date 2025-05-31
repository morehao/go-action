package main

import (
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

const (
	ModuleName = "LLM_call"
	LLMHost    = "https://api.deepseek.com/v1"
	LLMModel   = "deepseek-chat"
	LLMAPIKey  = "xxxx"
)

var llmClient = openai.NewClient(
	option.WithBaseURL(LLMHost),
	option.WithAPIKey(LLMAPIKey))

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Chat 普通调用
func Chat(ctx *gin.Context) {
	completion, err := llmClient.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("世界上最高的山峰"),
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
		Message: completion.Choices[0].Message.Content,
	})
}
