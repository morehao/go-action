package llmproxy

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// Client HTTP 客户端，用于调用 OpenAI 兼容的 API
type Client struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewClient 创建新的 HTTP 客户端
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		apiKey:  apiKey,
		client:  &http.Client{},
	}
}

// Chat 发送聊天请求
func (c *Client) Chat(ctx context.Context, req *types.ChatRequest) (*types.ChatResponse, error) {
	// 确保 Stream 为 false
	req.Stream = false

	// 序列化请求
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// 发送请求
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var chatResp types.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &chatResp, nil
}

// StreamChat 发送流式聊天请求
func (c *Client) StreamChat(ctx context.Context, req *types.ChatRequest) (<-chan types.StreamChunk, error) {
	// 确保 Stream 为 true
	req.Stream = true

	// 序列化请求
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "text/event-stream")
	if c.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// 发送请求
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	// 创建 channel
	ch := make(chan types.StreamChunk, 10)

	// 启动 goroutine 处理流式响应
	go func() {
		defer resp.Body.Close()
		defer close(ch)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			// 检查上下文是否已取消
			select {
			case <-ctx.Done():
				ch <- types.StreamChunk{Error: ctx.Err()}
				return
			default:
			}

			line := strings.TrimSpace(scanner.Text())

			// 跳过空行
			if line == "" {
				continue
			}

			// 检查是否是 SSE 格式的数据行
			if !strings.HasPrefix(line, "data: ") {
				continue
			}

			// 移除 "data: " 前缀
			data := strings.TrimPrefix(line, "data: ")

			// 检查是否是结束标记
			if data == "[DONE]" {
				return
			}

			// 解析 JSON
			var chunk types.StreamChunk
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				ch <- types.StreamChunk{Error: fmt.Errorf("decode chunk: %w", err)}
				continue
			}

			// 发送 chunk
			select {
			case ch <- chunk:
			case <-ctx.Done():
				ch <- types.StreamChunk{Error: ctx.Err()}
				return
			}
		}

		// 检查扫描错误
		if err := scanner.Err(); err != nil {
			ch <- types.StreamChunk{Error: fmt.Errorf("scan error: %w", err)}
		}
	}()

	return ch, nil
}
