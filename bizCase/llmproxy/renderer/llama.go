package renderer

import (
	"fmt"
	"strings"

	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// LlamaRenderer Llama 专用渲染器，使用 Python 函数签名格式
type LlamaRenderer struct{}

// RenderTools 实现 Renderer 接口
func (r *LlamaRenderer) RenderTools(req *types.ChatRequest) *types.ChatRequest {
	if len(req.Tools) == 0 {
		return req
	}

	// 创建请求副本
	newReq := *req
	newReq.Messages = make([]types.Message, len(req.Messages))
	copy(newReq.Messages, req.Messages)

	// 构建工具描述（Python 函数签名格式）
	var sb strings.Builder
	sb.WriteString("# Available Tools\n\n")
	sb.WriteString("Here is a list of tools that you have available to you:\n\n")

	for _, tool := range req.Tools {
		sb.WriteString("```python\n")
		sb.WriteString(fmt.Sprintf("def %s(", tool.Function.Name))

		// 构建参数列表
		if params, ok := tool.Function.Parameters["properties"].(map[string]interface{}); ok {
			first := true
			for name, prop := range params {
				if !first {
					sb.WriteString(", ")
				}
				first = false

				propType := "any"
				if propMap, ok := prop.(map[string]interface{}); ok {
					if t, ok := propMap["type"].(string); ok {
						propType = pythonType(t)
					}
				}

				sb.WriteString(fmt.Sprintf("%s: %s", name, propType))
			}
		}

		sb.WriteString(") -> dict:\n")
		sb.WriteString(fmt.Sprintf("    '''%s\n", tool.Function.Description))

		// 添加参数文档
		if params, ok := tool.Function.Parameters["properties"].(map[string]interface{}); ok && len(params) > 0 {
			sb.WriteString("\n    Args:\n")
			for name, prop := range params {
				propType := "any"
				propDesc := ""
				if propMap, ok := prop.(map[string]interface{}); ok {
					if t, ok := propMap["type"].(string); ok {
						propType = pythonType(t)
					}
					if d, ok := propMap["description"].(string); ok {
						propDesc = d
					}
				}
				sb.WriteString(fmt.Sprintf("        %s (%s): %s\n", name, propType, propDesc))
			}
		}

		sb.WriteString("    '''\n")
		sb.WriteString("    pass\n")
		sb.WriteString("```\n\n")
	}

	sb.WriteString("To call a function, respond with JSON in the following format:\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{"name": "function_name", "arguments": {"arg1": "value1", "arg2": "value2"}}`)
	sb.WriteString("\n```\n")

	toolDesc := sb.String()

	// 将工具描述添加到系统消息
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

// pythonType 将 JSON Schema 类型转换为 Python 类型
func pythonType(jsonType string) string {
	switch jsonType {
	case "string":
		return "str"
	case "number":
		return "float"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	case "array":
		return "list"
	case "object":
		return "dict"
	default:
		return "any"
	}
}
