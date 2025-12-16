package parser

import (
	"encoding/json"
	"regexp"
	"strings"
)

// XMLParser XML 标签解析器
type XMLParser struct {
	toolCallPattern *regexp.Regexp
}

// NewXMLParser 创建 XML 解析器
func NewXMLParser() *XMLParser {
	return &XMLParser{
		toolCallPattern: regexp.MustCompile(`<tool_call>([\s\S]*?)</tool_call>`),
	}
}

// Parse 实现 Parser 接口
func (p *XMLParser) Parse(content string) ([]InternalToolCall, string, error) {
	if p.toolCallPattern == nil {
		p.toolCallPattern = regexp.MustCompile(`<tool_call>([\s\S]*?)</tool_call>`)
	}

	var toolCalls []InternalToolCall
	var remainingContent strings.Builder

	// 查找所有 tool_call 标签
	matches := p.toolCallPattern.FindAllStringSubmatchIndex(content, -1)
	lastEnd := 0

	for _, match := range matches {
		// match[0] 是整个匹配的开始位置
		// match[1] 是整个匹配的结束位置
		// match[2] 是第一个捕获组的开始位置
		// match[3] 是第一个捕获组的结束位置

		// 添加标签之前的内容
		if match[0] > lastEnd {
			remainingContent.WriteString(content[lastEnd:match[0]])
		}

		// 提取工具调用内容
		toolCallJSON := strings.TrimSpace(content[match[2]:match[3]])

		// 解析工具调用
		var tc struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments"`
		}

		if err := json.Unmarshal([]byte(toolCallJSON), &tc); err != nil {
			// 解析失败，保留原内容
			remainingContent.WriteString(content[match[0]:match[1]])
		} else {
			toolCalls = append(toolCalls, InternalToolCall{
				Name:      tc.Name,
				Arguments: tc.Arguments,
			})
		}

		lastEnd = match[1]
	}

	// 添加剩余内容
	if lastEnd < len(content) {
		remainingContent.WriteString(content[lastEnd:])
	}

	return toolCalls, strings.TrimSpace(remainingContent.String()), nil
}

// StreamXMLParser 流式 XML 解析器
type StreamXMLParser struct {
	buffer    strings.Builder
	state     int // 0: 查找标签, 1: 解析工具调用, 2: 完成
	toolCalls []InternalToolCall
	tagStart  string
	tagEnd    string
}

// NewStreamXMLParser 创建流式 XML 解析器
func NewStreamXMLParser() *StreamXMLParser {
	return &StreamXMLParser{
		tagStart: "<tool_call>",
		tagEnd:   "</tool_call>",
	}
}

// Add 实现 StreamParser 接口
func (p *StreamXMLParser) Add(chunk string) ([]InternalToolCall, string, bool) {
	p.buffer.WriteString(chunk)
	bufStr := p.buffer.String()

	var output strings.Builder
	var newCalls []InternalToolCall

	for {
		if p.state == 0 {
			// 查找开始标签
			idx := strings.Index(bufStr, p.tagStart)
			if idx == -1 {
				// 没有找到开始标签，检查是否有部分匹配
				if len(bufStr) > len(p.tagStart) {
					// 保留可能的部分匹配
					keepLen := len(p.tagStart) - 1
					output.WriteString(bufStr[:len(bufStr)-keepLen])
					p.buffer.Reset()
					p.buffer.WriteString(bufStr[len(bufStr)-keepLen:])
				}
				break
			}

			// 找到开始标签
			output.WriteString(bufStr[:idx])
			bufStr = bufStr[idx+len(p.tagStart):]
			p.state = 1
		}

		if p.state == 1 {
			// 查找结束标签
			idx := strings.Index(bufStr, p.tagEnd)
			if idx == -1 {
				// 没有找到结束标签，继续等待
				p.buffer.Reset()
				p.buffer.WriteString(bufStr)
				break
			}

			// 找到结束标签，解析工具调用
			toolCallJSON := strings.TrimSpace(bufStr[:idx])
			var tc struct {
				Name      string                 `json:"name"`
				Arguments map[string]interface{} `json:"arguments"`
			}

			if err := json.Unmarshal([]byte(toolCallJSON), &tc); err == nil {
				call := InternalToolCall{
					Name:      tc.Name,
					Arguments: tc.Arguments,
				}
				p.toolCalls = append(p.toolCalls, call)
				newCalls = append(newCalls, call)
			}

			bufStr = bufStr[idx+len(p.tagEnd):]
			p.state = 0
		}
	}

	return newCalls, output.String(), false
}

// Flush 实现 StreamParser 接口
func (p *StreamXMLParser) Flush() ([]InternalToolCall, string) {
	remaining := p.buffer.String()
	p.buffer.Reset()
	return p.toolCalls, remaining
}
