package parser

import (
	"testing"
)

func TestXMLParser_Parse(t *testing.T) {
	parser := NewXMLParser()

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
			input:         `Let me check the weather. <tool_call>{"name": "get_weather", "arguments": {"city": "Beijing"}}</tool_call>`,
			wantToolCalls: 1,
			wantContent:   "Let me check the weather.",
			wantToolName:  "get_weather",
			wantArguments: map[string]interface{}{"city": "Beijing"},
		},
		{
			name:          "multiple tool calls",
			input:         `<tool_call>{"name": "tool1", "arguments": {"a": "1"}}</tool_call><tool_call>{"name": "tool2", "arguments": {"b": "2"}}</tool_call>`,
			wantToolCalls: 2,
			wantContent:   "",
		},
		{
			name:          "tool call with text before and after",
			input:         `Before <tool_call>{"name": "test", "arguments": {}}</tool_call> After`,
			wantToolCalls: 1,
			wantContent:   "Before  After",
			wantToolName:  "test",
		},
		{
			name:          "invalid json in tool call",
			input:         `<tool_call>{invalid json}</tool_call>`,
			wantToolCalls: 0,
			wantContent:   `<tool_call>{invalid json}</tool_call>`,
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

func TestStreamXMLParser_Add(t *testing.T) {
	parser := NewStreamXMLParser()

	// 测试流式解析
	chunks := []string{
		"Let me ",
		"check the weather. <tool_",
		"call>{\"name\": \"get_weather\", ",
		"\"arguments\": {\"city\": \"Beijing\"}}",
		"</tool_call> Done.",
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
		t.Errorf("StreamXMLParser got %d tool calls, want 1", len(allToolCalls))
	}

	if allToolCalls[0].Name != "get_weather" {
		t.Errorf("StreamXMLParser tool name = %q, want %q", allToolCalls[0].Name, "get_weather")
	}

	expectedContent := "Let me check the weather.  Done."
	if allContent != expectedContent {
		t.Errorf("StreamXMLParser content = %q, want %q", allContent, expectedContent)
	}
}
