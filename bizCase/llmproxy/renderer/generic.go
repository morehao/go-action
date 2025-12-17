/*
 * @Author: morehao morehao@qq.com
 * @Date: 2025-12-15 18:56:41
 * @LastEditors: morehao morehao@qq.com
 * @LastEditTime: 2025-12-16 18:57:59
 * @FilePath: /go-action/bizCase/llmproxy/renderer/generic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package renderer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// GenericRenderer 通用渲染器，使用 JSON 格式
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
	sb.WriteString("# 可用工具\n\n")
	sb.WriteString("你可以调用一个或多个函数来协助回答用户的问题。\n\n")
	sb.WriteString("## 工具列表\n\n")

	// 为每个工具生成详细描述
	for i, tool := range req.Tools {
		sb.WriteString(fmt.Sprintf("### %d. %s\n\n", i+1, tool.Function.Name))
		if tool.Function.Description != "" {
			sb.WriteString(fmt.Sprintf("**描述**: %s\n\n", tool.Function.Description))
		}

		// 输出参数详情
		if tool.Function.Parameters != nil {
			sb.WriteString("**参数**:\n")
			sb.WriteString("```json\n")
			paramsJSON, err := json.MarshalIndent(tool.Function.Parameters, "", "  ")
			if err == nil {
				sb.Write(paramsJSON)
			}
			sb.WriteString("\n```\n\n")
		}
	}

	sb.WriteString("## 完整工具定义\n\n")
	sb.WriteString("以下是 JSON 数组格式的完整工具定义：\n")
	sb.WriteString("```json\n")

	toolsJSON, err := json.MarshalIndent(req.Tools, "", "  ")
	if err == nil {
		sb.Write(toolsJSON)
	}
	// "即将使用<descrption> \n```json\n{\"name\": \"get_weather\", \"arguments\": \"{\\\"location\\\": \\\"西安\\\"}\"}\n```"

	sb.WriteString("\n```\n\n")
	sb.WriteString("## 调用格式\n\n")
	sb.WriteString("当你需要调用工具时，请返回以下格式的 函数调用信息总结 + JSON 对象：\n")
	sb.WriteString("<工具描述>\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{"name": "<工具名称>", "arguments": <参数JSON对象>}`)
	sb.WriteString("\n```\n\n")
	sb.WriteString("**注意事项**:\n")
	sb.WriteString("- 必须严格按照上述 JSON 格式返回\n")
	sb.WriteString("- `name` 字段必须是上述工具列表中的某个工具名称\n")
	sb.WriteString("- `arguments` 必须是符合该工具参数定义的 JSON 对象\n")
	sb.WriteString("- 你可以根据需要调用多个工具，每次调用返回一个独立的 JSON 对象\n")

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
