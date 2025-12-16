package parser

import (
	"encoding/json"
	"regexp"
	"strings"
)

// JSONParser JSON 格式解析器
type JSONParser struct {
	jsonPattern *regexp.Regexp
}

// NewJSONParser 创建 JSON 解析器
func NewJSONParser() *JSONParser {
	return &JSONParser{
		// 匹配包含 name 字段的 JSON 对象
		jsonPattern: regexp.MustCompile(`\{[^{}]*"name"\s*:\s*"[^"]+"[^{}]*\}`),
	}
}

// Parse 实现 Parser 接口
func (p *JSONParser) Parse(content string) ([]InternalToolCall, string, error) {
	if p.jsonPattern == nil {
		p.jsonPattern = regexp.MustCompile(`\{[^{}]*"name"\s*:\s*"[^"]+"[^{}]*\}`)
	}

	var toolCalls []InternalToolCall
	matches := p.jsonPattern.FindAllString(content, -1)

	for _, match := range matches {
		var tc struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments"`
		}

		if err := json.Unmarshal([]byte(match), &tc); err != nil {
			continue
		}

		// 只有当有 arguments 字段时才认为是工具调用
		if tc.Arguments != nil {
			toolCalls = append(toolCalls, InternalToolCall{
				Name:      tc.Name,
				Arguments: tc.Arguments,
			})
		}
	}

	// 移除已解析的 JSON
	remainingContent := content
	for _, match := range matches {
		remainingContent = strings.Replace(remainingContent, match, "", 1)
	}

	return toolCalls, strings.TrimSpace(remainingContent), nil
}

// StreamJSONParser 流式 JSON 解析器
type StreamJSONParser struct {
	buffer    strings.Builder
	toolCalls []InternalToolCall
	parser    *JSONParser
}

// NewStreamJSONParser 创建流式 JSON 解析器
func NewStreamJSONParser() *StreamJSONParser {
	return &StreamJSONParser{
		parser: NewJSONParser(),
	}
}

// Add 实现 StreamParser 接口
func (p *StreamJSONParser) Add(chunk string) ([]InternalToolCall, string, bool) {
	p.buffer.WriteString(chunk)
	bufStr := p.buffer.String()

	// 尝试解析
	calls, remaining, _ := p.parser.Parse(bufStr)

	var newCalls []InternalToolCall
	for _, call := range calls {
		// 检查是否是新的调用
		isNew := true
		for _, existing := range p.toolCalls {
			if existing.Name == call.Name {
				isNew = false
				break
			}
		}
		if isNew {
			p.toolCalls = append(p.toolCalls, call)
			newCalls = append(newCalls, call)
		}
	}

	// 更新缓冲区
	p.buffer.Reset()
	p.buffer.WriteString(remaining)

	// 如果没有剩余内容，返回空字符串
	output := ""
	if len(newCalls) == 0 {
		// 没有新的工具调用，可以输出内容
		if len(remaining) > 100 {
			// 保留一部分防止截断 JSON
			output = remaining[:len(remaining)-100]
			p.buffer.Reset()
			p.buffer.WriteString(remaining[len(remaining)-100:])
		}
	}

	return newCalls, output, false
}

// Flush 实现 StreamParser 接口
func (p *StreamJSONParser) Flush() ([]InternalToolCall, string) {
	remaining := p.buffer.String()
	p.buffer.Reset()
	return p.toolCalls, remaining
}
