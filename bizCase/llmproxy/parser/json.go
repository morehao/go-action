package parser

import (
	"encoding/json"
	"regexp"
	"strings"
)

// JSONParser JSON 格式解析器
type JSONParser struct {
	codeBlockPattern *regexp.Regexp
	jsonPattern      *regexp.Regexp
}

// NewJSONParser 创建 JSON 解析器
func NewJSONParser() *JSONParser {
	return &JSONParser{
		// 匹配 markdown 代码块中的 JSON
		codeBlockPattern: regexp.MustCompile("```(?:json)?\\s*([^`]+)```"),
		// 匹配包含 name 字段的 JSON 对象（支持嵌套）
		jsonPattern: regexp.MustCompile(`\{(?:[^{}]|\{[^{}]*\})*"name"\s*:\s*"[^"]+"(?:[^{}]|\{[^{}]*\})*\}`),
	}
}

// Parse 实现 Parser 接口
func (p *JSONParser) Parse(content string) ([]InternalToolCall, string, error) {
	if p.codeBlockPattern == nil {
		p.codeBlockPattern = regexp.MustCompile("```(?:json)?\\s*([^`]+)```")
	}

	var toolCalls []InternalToolCall
	originalContent := content
	remainingContent := content

	// 1. 先尝试从 markdown 代码块中提取 JSON
	codeBlockMatches := p.codeBlockPattern.FindAllStringSubmatch(content, -1)
	for _, match := range codeBlockMatches {
		if len(match) > 1 {
			jsonContent := strings.TrimSpace(match[1])
			// 尝试解析代码块中的 JSON
			if calls := p.parseJSON(jsonContent); len(calls) > 0 {
				toolCalls = append(toolCalls, calls...)
				// 从剩余内容中移除整个代码块
				remainingContent = strings.Replace(remainingContent, match[0], "", 1)
			}
		}
	}

	// 2. 如果没有找到代码块，尝试使用更智能的 JSON 提取
	if len(toolCalls) == 0 {
		// 使用括号匹配来提取完整的 JSON 对象
		matches := p.extractJSONObjects(content)
		for _, match := range matches {
			if calls := p.parseJSON(match); len(calls) > 0 {
				toolCalls = append(toolCalls, calls...)
				remainingContent = strings.Replace(remainingContent, match, "", 1)
			}
		}
	}

	// 如果成功解析到工具调用，返回剩余内容；否则返回原始内容
	if len(toolCalls) > 0 {
		return toolCalls, strings.TrimSpace(remainingContent), nil
	}

	return toolCalls, strings.TrimSpace(originalContent), nil
}

// extractJSONObjects 从文本中提取所有可能的 JSON 对象（通过括号匹配）
func (p *JSONParser) extractJSONObjects(content string) []string {
	var objects []string
	var start int = -1
	var depth int

	for i := 0; i < len(content); i++ {
		if content[i] == '{' {
			if depth == 0 {
				start = i
			}
			depth++
		} else if content[i] == '}' {
			depth--
			if depth == 0 && start >= 0 {
				// 找到一个完整的 JSON 对象
				jsonStr := content[start : i+1]
				// 检查是否包含 "name" 字段
				if strings.Contains(jsonStr, `"name"`) {
					objects = append(objects, jsonStr)
				}
				start = -1
			}
		}
	}

	return objects
}

// parseJSON 解析 JSON 字符串，提取工具调用
func (p *JSONParser) parseJSON(jsonStr string) []InternalToolCall {
	var calls []InternalToolCall

	// 定义一个临时结构体，arguments 可以是 string 或 map
	var tc struct {
		Name      string      `json:"name"`
		Arguments interface{} `json:"arguments"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &tc); err != nil {
		return calls
	}

	// 处理 arguments
	var args map[string]interface{}
	switch v := tc.Arguments.(type) {
	case string:
		// arguments 是字符串，尝试解析为 map
		if err := json.Unmarshal([]byte(v), &args); err != nil {
			// 如果解析失败，返回空
			return calls
		}
	case map[string]interface{}:
		// arguments 本身就是 map
		args = v
	default:
		// 不支持的类型
		return calls
	}

	// 只有当有 arguments 时才认为是工具调用
	if len(args) > 0 {
		calls = append(calls, InternalToolCall{
			Name:      tc.Name,
			Arguments: args,
		})
	}

	return calls
}

// StreamJSONParser 流式 JSON 解析器
type StreamJSONParser struct {
	buffer            strings.Builder
	toolCalls         []InternalToolCall
	parser            *JSONParser
	processedJSONStrs map[string]bool // 记录已处理的 JSON 字符串，避免重复
}

// NewStreamJSONParser 创建流式 JSON 解析器
func NewStreamJSONParser() *StreamJSONParser {
	return &StreamJSONParser{
		parser:            NewJSONParser(),
		processedJSONStrs: make(map[string]bool),
	}
}

// Add 实现 StreamParser 接口
func (p *StreamJSONParser) Add(chunk string) ([]InternalToolCall, string, bool) {
	p.buffer.WriteString(chunk)
	bufStr := p.buffer.String()

	// 尝试解析
	calls, remaining, _ := p.parser.Parse(bufStr)

	var newCalls []InternalToolCall
	hasNewCall := false

	// 检查是否有新的工具调用
	for _, call := range calls {
		// 将工具调用序列化为字符串，用于去重
		callKey := p.getCallKey(call)
		if !p.processedJSONStrs[callKey] {
			p.processedJSONStrs[callKey] = true
			p.toolCalls = append(p.toolCalls, call)
			newCalls = append(newCalls, call)
			hasNewCall = true
		}
	}

	// 更新缓冲区
	if hasNewCall {
		// 如果发现新的工具调用，更新缓冲区为剩余内容
		p.buffer.Reset()
		p.buffer.WriteString(remaining)
		return newCalls, "", false
	}

	// 如果没有新的工具调用，但缓冲区很大，可以输出一些安全的内容
	if len(bufStr) > 200 {
		// 保留最后 200 个字符，防止截断正在形成的 JSON
		output := bufStr[:len(bufStr)-200]
		p.buffer.Reset()
		p.buffer.WriteString(bufStr[len(bufStr)-200:])
		return nil, output, false
	}

	return nil, "", false
}

// getCallKey 生成工具调用的唯一键
func (p *StreamJSONParser) getCallKey(call InternalToolCall) string {
	// 简单地使用 JSON 序列化作为键
	data, _ := json.Marshal(call)
	return string(data)
}

// Flush 实现 StreamParser 接口
func (p *StreamJSONParser) Flush() ([]InternalToolCall, string) {
	remaining := p.buffer.String()
	p.buffer.Reset()
	// 不返回已收集的工具调用，因为它们已经在 Add() 中返回过了
	// 只返回剩余的文本内容
	return nil, remaining
}
