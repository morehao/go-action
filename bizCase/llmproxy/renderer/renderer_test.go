package renderer

import (
	"strings"
	"testing"

	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

func TestGenericRenderer_RenderTools(t *testing.T) {
	renderer := &GenericRenderer{}

	tool := types.Tool{
		Type: "function",
		Function: types.ToolFunction{
			Name:        "get_weather",
			Description: "Get weather information",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"city": map[string]interface{}{
						"type":        "string",
						"description": "City name",
					},
				},
			},
		},
	}

	req := &types.ChatRequest{
		Model: "test-model",
		Messages: []types.Message{
			{Role: "user", Content: "What's the weather?"},
		},
		Tools: []types.Tool{tool},
	}

	result := renderer.RenderTools(req)

	// 验证系统消息被添加
	if len(result.Messages) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(result.Messages))
	}

	if result.Messages[0].Role != "system" {
		t.Errorf("Expected first message role to be system, got %s", result.Messages[0].Role)
	}

	// 验证工具描述包含关键信息
	content := result.Messages[0].Content
	if !strings.Contains(content, "Available Tools") {
		t.Error("System message should contain 'Available Tools'")
	}

	if !strings.Contains(content, "<tools>") {
		t.Error("System message should contain '<tools>' tag")
	}

	if !strings.Contains(content, "get_weather") {
		t.Error("System message should contain tool name 'get_weather'")
	}

	if !strings.Contains(content, "<tool_call>") {
		t.Error("System message should contain '<tool_call>' example")
	}
}

func TestQwenRenderer_RenderTools(t *testing.T) {
	renderer := &QwenRenderer{}

	tool := types.Tool{
		Type: "function",
		Function: types.ToolFunction{
			Name:        "calculate",
			Description: "Perform calculation",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"expression": map[string]interface{}{
						"type":        "string",
						"description": "Math expression",
					},
				},
			},
		},
	}

	req := &types.ChatRequest{
		Model: "qwen",
		Messages: []types.Message{
			{Role: "user", Content: "Calculate 2+2"},
		},
		Tools: []types.Tool{tool},
	}

	result := renderer.RenderTools(req)

	content := result.Messages[0].Content

	// 验证 Qwen 特定格式
	if !strings.Contains(content, "<function>") {
		t.Error("Qwen format should contain '<function>' tag")
	}

	if !strings.Contains(content, "<name>calculate</name>") {
		t.Error("Qwen format should contain tool name in tags")
	}

	if !strings.Contains(content, "<parameters>") {
		t.Error("Qwen format should contain '<parameters>' tag")
	}

	if !strings.Contains(content, "<IMPORTANT>") {
		t.Error("Qwen format should contain '<IMPORTANT>' reminder")
	}
}

func TestLlamaRenderer_RenderTools(t *testing.T) {
	renderer := &LlamaRenderer{}

	tool := types.Tool{
		Type: "function",
		Function: types.ToolFunction{
			Name:        "search",
			Description: "Search the web",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"query": map[string]interface{}{
						"type":        "string",
						"description": "Search query",
					},
				},
			},
		},
	}

	req := &types.ChatRequest{
		Model: "llama",
		Messages: []types.Message{
			{Role: "user", Content: "Search for Go language"},
		},
		Tools: []types.Tool{tool},
	}

	result := renderer.RenderTools(req)

	content := result.Messages[0].Content

	// 验证 Python 函数签名格式
	if !strings.Contains(content, "```python") {
		t.Error("Llama format should contain Python code block")
	}

	if !strings.Contains(content, "def search(") {
		t.Error("Llama format should contain function definition")
	}

	if !strings.Contains(content, "query: str") {
		t.Error("Llama format should contain typed parameters")
	}

	if !strings.Contains(content, "Args:") {
		t.Error("Llama format should contain docstring with Args")
	}
}

func TestNewRenderer(t *testing.T) {
	tests := []struct {
		format   string
		wantType string
	}{
		{"qwen", "*renderer.QwenRenderer"},
		{"llama", "*renderer.LlamaRenderer"},
		{"generic", "*renderer.GenericRenderer"},
		{"unknown", "*renderer.GenericRenderer"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			r := NewRenderer(tt.format)
			if r == nil {
				t.Error("NewRenderer should not return nil")
			}
		})
	}
}
