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
	sb.WriteString("**重要**：当你需要调用工具时，必须按照以下两步骤格式返回，缺一不可：\n\n")
	sb.WriteString("**步骤1**：在第一行写一句话，描述你要调用什么工具以及传入什么参数\n\n")
	sb.WriteString("**步骤2**：空一行后，写 JSON 代码块（以 ```json 开头，以 ``` 结尾）\n\n")
	sb.WriteString("### 标准示例（必须遵循）\n\n")
	sb.WriteString("调用天气查询工具获取西安的天气信息\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{"name": "get_weather", "arguments": "{\"location\": \"西安\"}"}`)
	sb.WriteString("\n```\n\n")
	sb.WriteString("### 错误示例（禁止这样做）\n\n")
	sb.WriteString("禁止直接返回 JSON 代码块而没有前面的描述：\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{"name": "get_weather", "arguments": "{\"location\": \"西安\"}"}`)
	sb.WriteString("\n```\n\n")
	sb.WriteString("### JSON 格式说明\n\n")
	sb.WriteString("- `name`: 工具名称（必须是上述工具列表中的某个）\n")
	sb.WriteString("- `arguments`: **必须是转义后的 JSON 字符串**（注意：是字符串，不是对象）\n")
	sb.WriteString("- 在 `arguments` 字符串中，所有双引号必须使用 `\\\"` 转义\n\n")
	sb.WriteString("### 完整格式模板\n\n")
	sb.WriteString("<一句话描述调用的工具和参数>\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{"name": "<工具名称>", "arguments": "<转义后的JSON字符串>"}`)
	sb.WriteString("\n```\n\n")
	sb.WriteString("**记住**：每次调用工具时，必须同时包含描述文字和 JSON 代码块！\n")

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
