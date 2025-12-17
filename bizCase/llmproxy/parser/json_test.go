package parser

import (
	"testing"
)

func TestJSONParser_Parse(t *testing.T) {
	parser := NewJSONParser()

	tests := []struct {
		name          string
		input         string
		wantToolCalls int
		wantContent   string
		wantToolName  string
		wantArguments map[string]interface{}
	}{
		{
			name:          "no tool call",
			input:         "This is just regular text",
			wantToolCalls: 0,
			wantContent:   "This is just regular text",
		},
		{
			name:          "single tool call",
			input:         `Let me check the weather. {"name": "get_weather", "arguments": {"city": "Beijing"}}`,
			wantToolCalls: 1,
			wantContent:   "Let me check the weather.",
			wantToolName:  "get_weather",
			wantArguments: map[string]interface{}{"city": "Beijing"},
		},
		{
			name:          "multiple tool calls",
			input:         `{"name": "tool1", "arguments": {"a": "1"}}{"name": "tool2", "arguments": {"b": "2"}}`,
			wantToolCalls: 2,
			wantContent:   "",
		},
		{
			name:          "tool call with text before and after",
			input:         `Before {"name": "test", "arguments": {"key": "value"}} After`,
			wantToolCalls: 1,
			wantContent:   "Before  After",
			wantToolName:  "test",
		},
		{
			name:          "invalid json",
			input:         `{invalid json}`,
			wantToolCalls: 0,
			wantContent:   `{invalid json}`,
		},
		{
			name:          "json without name field",
			input:         `{"arguments": {"city": "Shanghai"}}`,
			wantToolCalls: 0,
			wantContent:   `{"arguments": {"city": "Shanghai"}}`,
		},
		{
			name:          "json with name but without arguments",
			input:         `{"name": "get_weather"}`,
			wantToolCalls: 0,
			wantContent:   `{"name": "get_weather"}`,
		},
		{
			name:          "nested json objects",
			input:         `Call: {"name": "complex_tool", "arguments": {"nested": {"key": "value"}, "array": [1, 2, 3]}}`,
			wantToolCalls: 1,
			wantContent:   "Call:",
			wantToolName:  "complex_tool",
		},
		{
			name:          "markdown code block with json",
			input:         "即将使用获取指定位置的天气\n```json\n{\"name\": \"get_weather\", \"arguments\": {\"location\": \"西安\"}}\n```",
			wantToolCalls: 1,
			wantContent:   "即将使用获取指定位置的天气",
			wantToolName:  "get_weather",
			wantArguments: map[string]interface{}{"location": "西安"},
		},
		{
			name:          "markdown code block with string arguments",
			input:         "即将使用获取指定位置的天气\n```json\n{\"name\": \"get_weather\", \"arguments\": \"{\\\"location\\\": \\\"西安\\\"}\"}\n```",
			wantToolCalls: 1,
			wantContent:   "即将使用获取指定位置的天气",
			wantToolName:  "get_weather",
			wantArguments: map[string]interface{}{"location": "西安"},
		},
		{
			name:          "markdown code block without language tag",
			input:         "让我查询一下\n```\n{\"name\": \"search\", \"arguments\": {\"query\": \"test\"}}\n```\n结果如下",
			wantToolCalls: 1,
			wantContent:   "让我查询一下\n\n结果如下",
			wantToolName:  "search",
			wantArguments: map[string]interface{}{"query": "test"},
		},
		{
			name:          "pure json with string arguments",
			input:         `{"name": "get_weather", "arguments": "{\"location\": \"北京\"}"}`,
			wantToolCalls: 1,
			wantContent:   "",
			wantToolName:  "get_weather",
			wantArguments: map[string]interface{}{"location": "北京"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toolCalls, content, err := parser.Parse(tt.input)
			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}

			if len(toolCalls) != tt.wantToolCalls {
				t.Errorf("Parse() got %d tool calls, want %d", len(toolCalls), tt.wantToolCalls)
			}

			if content != tt.wantContent {
				t.Errorf("Parse() content = %q, want %q", content, tt.wantContent)
			}

			if tt.wantToolCalls > 0 && tt.wantToolName != "" {
				if toolCalls[0].Name != tt.wantToolName {
					t.Errorf("Parse() tool name = %q, want %q", toolCalls[0].Name, tt.wantToolName)
				}

				if tt.wantArguments != nil {
					for key, wantVal := range tt.wantArguments {
						if gotVal, ok := toolCalls[0].Arguments[key]; !ok || gotVal != wantVal {
							t.Errorf("Parse() argument %q = %v, want %v", key, gotVal, wantVal)
						}
					}
				}
			}
		})
	}
}

func TestStreamJSONParser_Add(t *testing.T) {
	parser := NewStreamJSONParser()

	// 测试流式解析
	chunks := []string{
		"Let me ",
		"check the weather. {\"name\": \"get_weather\", ",
		"\"arguments\": {\"city\": \"Beijing\"}}",
		" Done.",
	}

	var allToolCalls []InternalToolCall
	var allContent string

	for _, chunk := range chunks {
		toolCalls, content, _ := parser.Add(chunk)
		allToolCalls = append(allToolCalls, toolCalls...)
		allContent += content
	}

	// 刷新缓冲区
	finalToolCalls, finalContent := parser.Flush()
	allToolCalls = append(allToolCalls, finalToolCalls...)
	allContent += finalContent

	if len(allToolCalls) != 1 {
		t.Errorf("StreamJSONParser got %d tool calls, want 1", len(allToolCalls))
	}

	if allToolCalls[0].Name != "get_weather" {
		t.Errorf("StreamJSONParser tool name = %q, want %q", allToolCalls[0].Name, "get_weather")
	}

	expectedContent := "Let me check the weather. Done."
	if allContent != expectedContent {
		t.Errorf("StreamJSONParser content = %q, want %q", allContent, expectedContent)
	}
}

func TestStreamJSONParser_MultipleToolCalls(t *testing.T) {
	parser := NewStreamJSONParser()

	// 测试多个工具调用的流式解析
	chunks := []string{
		`First call: {"name": "tool1", "arguments": {"param1": "value1"}}`,
		` Second call: {"name": "tool2", "arguments": {"param2": "value2"}}`,
	}

	var allToolCalls []InternalToolCall

	for _, chunk := range chunks {
		toolCalls, _, _ := parser.Add(chunk)
		allToolCalls = append(allToolCalls, toolCalls...)
	}

	// 刷新缓冲区
	finalToolCalls, _ := parser.Flush()
	allToolCalls = append(allToolCalls, finalToolCalls...)

	if len(allToolCalls) != 2 {
		t.Errorf("StreamJSONParser got %d tool calls, want 2", len(allToolCalls))
	}

	if allToolCalls[0].Name != "tool1" {
		t.Errorf("StreamJSONParser first tool name = %q, want %q", allToolCalls[0].Name, "tool1")
	}

	if allToolCalls[1].Name != "tool2" {
		t.Errorf("StreamJSONParser second tool name = %q, want %q", allToolCalls[1].Name, "tool2")
	}
}

func TestStreamJSONParser_PartialJSON(t *testing.T) {
	parser := NewStreamJSONParser()

	// 测试部分 JSON 的处理
	chunks := []string{
		`{"na`,
		`me": "test_tool", "arg`,
		`uments": {"key": "val`,
		`ue"}}`,
	}

	var allToolCalls []InternalToolCall

	for _, chunk := range chunks {
		toolCalls, _, _ := parser.Add(chunk)
		allToolCalls = append(allToolCalls, toolCalls...)
	}

	// 刷新缓冲区
	finalToolCalls, _ := parser.Flush()
	allToolCalls = append(allToolCalls, finalToolCalls...)

	if len(allToolCalls) != 1 {
		t.Errorf("StreamJSONParser got %d tool calls, want 1", len(allToolCalls))
	}

	if allToolCalls[0].Name != "test_tool" {
		t.Errorf("StreamJSONParser tool name = %q, want %q", allToolCalls[0].Name, "test_tool")
	}

	if val, ok := allToolCalls[0].Arguments["key"]; !ok || val != "value" {
		t.Errorf("StreamJSONParser argument key = %v, want %v", val, "value")
	}
}
