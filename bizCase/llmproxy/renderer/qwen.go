/*
 * @Author: morehao morehao@qq.com
 * @Date: 2025-12-15 18:56:41
 * @LastEditors: morehao morehao@qq.com
 * @LastEditTime: 2025-12-16 14:37:55
 * @FilePath: /go-action/bizCase/llmproxy/renderer/qwen.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package renderer

import (
	"fmt"
	"strings"

	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// QwenRenderer Qwen 专用渲染器
type QwenRenderer struct{}

// RenderTools 实现 Renderer 接口
func (r *QwenRenderer) RenderTools(req *types.ChatRequest) *types.ChatRequest {
	if len(req.Tools) == 0 {
		return req
	}

	// 创建请求副本
	newReq := *req
	newReq.Messages = make([]types.Message, len(req.Messages))
	copy(newReq.Messages, req.Messages)

	// 构建工具描述（Qwen 特定格式）
	var sb strings.Builder
	sb.WriteString("# Tools\n\n")
	sb.WriteString("You have access to the following functions:\n\n")
	sb.WriteString("<tools>")

	for _, tool := range req.Tools {
		sb.WriteString("\n<function>\n")
		sb.WriteString(fmt.Sprintf("<name>%s</name>\n", tool.Function.Name))
		if tool.Function.Description != "" {
			sb.WriteString(fmt.Sprintf("<description>%s</description>\n", tool.Function.Description))
		}
		sb.WriteString("<parameters>")

		// 解析参数
		if params, ok := tool.Function.Parameters["properties"].(map[string]interface{}); ok {
			for name, prop := range params {
				sb.WriteString("\n<parameter>")
				sb.WriteString(fmt.Sprintf("\n<name>%s</name>", name))

				if propMap, ok := prop.(map[string]interface{}); ok {
					if propType, ok := propMap["type"].(string); ok {
						sb.WriteString(fmt.Sprintf("\n<type>%s</type>", propType))
					}
					if desc, ok := propMap["description"].(string); ok {
						sb.WriteString(fmt.Sprintf("\n<description>%s</description>", desc))
					}
				}

				sb.WriteString("\n</parameter>")
			}
		}

		sb.WriteString("\n</parameters>")
		sb.WriteString("\n</function>")
	}

	sb.WriteString("\n</tools>\n\n")
	sb.WriteString("If you choose to call a function ONLY reply in the following format with NO suffix:\n\n")
	sb.WriteString("<tool_call>\n")
	sb.WriteString("<function=example_function_name>\n")
	sb.WriteString("<parameter=example_parameter_1>\nvalue_1\n</parameter>\n")
	sb.WriteString("<parameter=example_parameter_2>\nvalue_2\n</parameter>\n")
	sb.WriteString("</function>\n</tool_call>\n\n")
	sb.WriteString("<IMPORTANT>\n")
	sb.WriteString("Reminder:\n")
	sb.WriteString("- Function calls MUST follow the specified format\n")
	sb.WriteString("- Required parameters MUST be specified\n")
	sb.WriteString("- You may provide optional reasoning for your function call in natural language BEFORE the function call, but NOT after\n")
	sb.WriteString("- If there is no function call available, answer the question like normal with your current knowledge and do not tell the user about function calls\n")
	sb.WriteString("</IMPORTANT>")

	toolDesc := sb.String()

	// 将工具描述添加到系统消息
	if len(newReq.Messages) > 0 && newReq.Messages[0].Role == "system" {
		newReq.Messages[0].Content = fmt.Sprintf("%s\n\n%s", newReq.Messages[0].Content, toolDesc)
	} else {
		systemMsg := types.Message{
			Role:    "system",
			Content: "You are Qwen, a helpful AI assistant that can interact with functions to solve tasks.\n\n" + toolDesc,
		}
		newReq.Messages = append([]types.Message{systemMsg}, newReq.Messages...)
	}

	return &newReq
}
