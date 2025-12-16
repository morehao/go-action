package renderer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// GenericRenderer 通用渲染器，使用 XML 格式
type GenericRenderer struct{}

// RenderTools 实现 Renderer 接口
func (r *GenericRenderer) RenderTools(req *types.ChatRequest) *types.ChatRequest {
	if len(req.Tools) == 0 {
		return req
	}

	// 创建请求副本
	newReq := *req
	newReq.Messages = make([]types.Message, len(req.Messages))
	copy(newReq.Messages, req.Messages)

	// 构建工具描述
	var sb strings.Builder
	sb.WriteString("# Available Tools\n\n")
	sb.WriteString("You may call one or more functions to assist with the user query.\n\n")
	sb.WriteString("You are provided with function signatures within <tools></tools> XML tags:\n")
	sb.WriteString("<tools>")

	for _, tool := range req.Tools {
		sb.WriteString("\n")
		toolJSON, err := json.MarshalIndent(tool, "", "  ")
		if err != nil {
			continue
		}
		sb.Write(toolJSON)
	}

	sb.WriteString("\n</tools>\n\n")
	sb.WriteString("For each function call, return a json object with function name and arguments within <tool_call></tool_call> XML tags:\n")
	sb.WriteString("<tool_call>\n")
	sb.WriteString(`{"name": "<function-name>", "arguments": <args-json-object>}`)
	sb.WriteString("\n</tool_call>\n")

	toolDesc := sb.String()

	// 将工具描述添加到系统消息或创建新的系统消息
	if len(newReq.Messages) > 0 && newReq.Messages[0].Role == "system" {
		newReq.Messages[0].Content = fmt.Sprintf("%s\n\n%s", newReq.Messages[0].Content, toolDesc)
	} else {
		systemMsg := types.Message{
			Role:    "system",
			Content: toolDesc,
		}
		newReq.Messages = append([]types.Message{systemMsg}, newReq.Messages...)
	}

	return &newReq
}
