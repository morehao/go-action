package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	"github.com/openai/openai-go"
)

func FunctionCall(ctx *gin.Context) {
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("西安今天天气怎么样"),
		},
		Tools: []openai.ChatCompletionToolParam{
			{
				Function: openai.FunctionDefinitionParam{
					Name:        "get_weather",
					Description: openai.String("获取指定位置的天气"),
					Parameters: openai.FunctionParameters{
						"type": "object",
						"properties": map[string]interface{}{
							"location": map[string]string{
								"type": "string",
							},
						},
						"required": []string{"location"},
					},
				},
			},
		},
		Seed:  openai.Int(0),
		Model: LLMModel,
	}

	completion, err := llmClient.Chat.Completions.New(ctx, params)
	if err != nil {
		ctx.JSON(200, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	toolCalls := completion.Choices[0].Message.ToolCalls

	if len(toolCalls) == 0 {
		glog.Warn(ctx, "[FunctionCall] No tool calls found")
		ctx.JSON(200, Response{
			Success: true,
			Message: "暂未找到天气数据",
		})
		return
	}

	glog.Infof(ctx, "[FunctionCall] Tool calls: %s", glog.ToJsonString(toolCalls))

	params.Messages = append(params.Messages, completion.Choices[0].Message.ToParam())
	for _, toolCall := range toolCalls {
		if toolCall.Function.Name == "get_weather" {
			// 从参数中提取位置
			var args map[string]interface{}
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			if err != nil {
				glog.Errorf(ctx, "[FunctionCall] Failed to parse arguments: %s", err.Error())
				ctx.JSON(200, Response{
					Success: false,
					Message: err.Error(),
				})
				return
			}
			location := args["location"].(string)

			// 获取天气数据
			weatherData := getWeather(location)
			glog.Infof(ctx, "[FunctionCall] Weather in %s: %s", location, weatherData)
			params.Messages = append(params.Messages, openai.ToolMessage(weatherData, toolCall.ID))
		}
	}

	newCompletion, newChatErr := llmClient.Chat.Completions.New(ctx, params)
	if newChatErr != nil {
		ctx.JSON(200, Response{
			Success: false,
			Message: newChatErr.Error(),
		})
		return
	}

	ctx.JSON(200, Response{
		Success: true,
		Message: newCompletion.Choices[0].Message.Content,
	})
}

// Mock天气数据
func getWeather(location string) string {
	return fmt.Sprintf("今天%s天气晴朗，温度%d摄氏度", location, 30)
}
