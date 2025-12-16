package llmproxy

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/morehao/go-action/bizCase/llmproxy/parser"
)

func TestToOpenAIToolCalls(t *testing.T) {
	internalCalls := []parser.InternalToolCall{
		{
			Name: "get_weather",
			Arguments: map[string]interface{}{
				"city": "Beijing",
				"unit": "celsius",
			},
		},
		{
			Name: "calculate",
			Arguments: map[string]interface{}{
				"expression": "2+2",
			},
		},
	}

	result := ToOpenAIToolCalls(internalCalls)

	if len(result) != 2 {
		t.Errorf("Expected 2 tool calls, got %d", len(result))
	}

	// 验证第一个工具调用
	if result[0].Type != "function" {
		t.Errorf("Expected type 'function', got %s", result[0].Type)
	}

	if result[0].Function.Name != "get_weather" {
		t.Errorf("Expected name 'get_weather', got %s", result[0].Function.Name)
	}

	if result[0].ID == "" {
		t.Error("ID should not be empty")
	}

	// 验证参数 JSON
	var args1 map[string]interface{}
	if err := json.Unmarshal([]byte(result[0].Function.Arguments), &args1); err != nil {
		t.Errorf("Failed to unmarshal arguments: %v", err)
	}

	if args1["city"] != "Beijing" {
		t.Errorf("Expected city 'Beijing', got %v", args1["city"])
	}

	// 验证第二个工具调用
	if result[1].Function.Name != "calculate" {
		t.Errorf("Expected name 'calculate', got %s", result[1].Function.Name)
	}
}

func TestChatRequestJSON(t *testing.T) {
	req := &ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "user", Content: "Hello"},
		},
		Tools: []Tool{
			{
				Type: "function",
				Function: ToolFunction{
					Name:        "test_function",
					Description: "A test function",
					Parameters: map[string]interface{}{
						"type": "object",
					},
				},
			},
		},
		Stream: false,
	}

	// 测试 JSON 序列化
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	// 测试 JSON 反序列化
	var decoded ChatRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal request: %v", err)
	}

	if decoded.Model != req.Model {
		t.Errorf("Model mismatch: got %s, want %s", decoded.Model, req.Model)
	}

	if len(decoded.Messages) != len(req.Messages) {
		t.Errorf("Messages length mismatch: got %d, want %d", len(decoded.Messages), len(req.Messages))
	}

	if len(decoded.Tools) != len(req.Tools) {
		t.Errorf("Tools length mismatch: got %d, want %d", len(decoded.Tools), len(req.Tools))
	}
}

func TestChatResponseJSON(t *testing.T) {
	resp := &ChatResponse{
		ID:      "chatcmpl-123",
		Object:  "chat.completion",
		Created: 1234567890,
		Model:   "gpt-3.5-turbo",
		Choices: []Choice{
			{
				Index: 0,
				Message: Message{
					Role:    "assistant",
					Content: "Hello!",
					ToolCalls: []ToolCall{
						{
							ID:   "call_123",
							Type: "function",
							Function: ToolCallFunction{
								Name:      "test",
								Arguments: `{"arg":"value"}`,
							},
						},
					},
				},
				FinishReason: "tool_calls",
			},
		},
		Usage: &Usage{
			PromptTokens:     10,
			CompletionTokens: 20,
			TotalTokens:      30,
		},
	}

	// 测试 JSON 序列化
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Failed to marshal response: %v", err)
	}

	// 测试 JSON 反序列化
	var decoded ChatResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if decoded.ID != resp.ID {
		t.Errorf("ID mismatch: got %s, want %s", decoded.ID, resp.ID)
	}

	if len(decoded.Choices) != len(resp.Choices) {
		t.Errorf("Choices length mismatch: got %d, want %d", len(decoded.Choices), len(resp.Choices))
	}

	if decoded.Usage.TotalTokens != resp.Usage.TotalTokens {
		t.Errorf("Total tokens mismatch: got %d, want %d", decoded.Usage.TotalTokens, resp.Usage.TotalTokens)
	}
}

func TestGenerateToolCallID(t *testing.T) {
	id1 := generateToolCallID()
	id2 := generateToolCallID()

	if id1 == "" {
		t.Error("Generated ID should not be empty")
	}

	if !strings.HasPrefix(id1, "call_") {
		t.Errorf("Generated ID should start with 'call_', got %s", id1)
	}

	if id1 == id2 {
		t.Error("Generated IDs should be unique")
	}

	if len(id1) != 29 { // "call_" (5) + 24 characters
		t.Errorf("Expected ID length 29, got %d", len(id1))
	}
}
