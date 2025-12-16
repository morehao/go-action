/*
 * @Author: morehao morehao@qq.com
 * @Date: 2025-12-15 18:56:41
 * @LastEditors: morehao morehao@qq.com
 * @LastEditTime: 2025-12-16 14:37:47
 * @FilePath: /go-action/bizCase/llmproxy/types.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package llmproxy

import (
	"encoding/json"
	"time"

	"github.com/morehao/go-action/bizCase/llmproxy/parser"
	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// 类型别名，保持向后兼容
type ChatRequest = types.ChatRequest
type ChatResponse = types.ChatResponse
type Choice = types.Choice
type Message = types.Message
type Tool = types.Tool
type ToolFunction = types.ToolFunction
type ToolCall = types.ToolCall
type ToolCallFunction = types.ToolCallFunction
type Usage = types.Usage
type StreamChunk = types.StreamChunk
type ChunkChoice = types.ChunkChoice

// ToOpenAIToolCalls 将内部工具调用转换为 OpenAI 格式
func ToOpenAIToolCalls(calls []parser.InternalToolCall) []types.ToolCall {
	result := make([]types.ToolCall, len(calls))
	for i, tc := range calls {
		args, _ := json.Marshal(tc.Arguments)
		result[i] = types.ToolCall{
			ID:   generateToolCallID(),
			Type: "function",
			Function: types.ToolCallFunction{
				Name:      tc.Name,
				Arguments: string(args),
			},
		}
	}
	return result
}

// generateToolCallID 生成工具调用 ID
func generateToolCallID() string {
	return "call_" + randomString(24)
}

// randomString 生成随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
