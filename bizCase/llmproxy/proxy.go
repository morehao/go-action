package llmproxy

import (
	"context"
	"fmt"
	"time"

	"github.com/morehao/go-action/bizCase/llmproxy/parser"
	"github.com/morehao/go-action/bizCase/llmproxy/renderer"
	"github.com/morehao/go-action/bizCase/llmproxy/types"
)

// Proxy 代理客户端
type Proxy struct {
	client   *Client
	renderer renderer.Renderer
	parser   parser.Parser
}

// NewProxy 创建新的代理实例
func NewProxy(client *Client, format string) *Proxy {
	return &Proxy{
		client:   client,
		renderer: renderer.NewRenderer(format),
		parser:   parser.NewParser(format),
	}
}

// Chat 处理聊天请求
// supportsNative: 模型是否支持原生 function call
func (p *Proxy) Chat(ctx context.Context, req *types.ChatRequest, supportsNative bool) (*types.ChatResponse, error) {
	// 如果模型支持原生 function call 或没有工具，直接转发
	if supportsNative || len(req.Tools) == 0 {
		return p.client.Chat(ctx, req)
	}

	// 不支持原生 function call，需要模拟
	return p.handleSimulatedTools(ctx, req)
}

// handleSimulatedTools 处理模拟的工具调用
func (p *Proxy) handleSimulatedTools(ctx context.Context, req *types.ChatRequest) (*types.ChatResponse, error) {
	// 1. 使用 renderer 将工具定义转换为提示词
	modifiedReq := p.renderer.RenderTools(req)

	// 2. 调用模型
	resp, err := p.client.Chat(ctx, modifiedReq)
	if err != nil {
		return nil, fmt.Errorf("call model: %w", err)
	}

	// 3. 使用 parser 解析工具调用
	if len(resp.Choices) > 0 {
		toolCalls, content, err := p.parser.Parse(resp.Choices[0].Message.Content)
		if err != nil {
			// 解析失败，返回原始响应
			return resp, nil
		}

		// 4. 转换为 OpenAI 格式
		resp.Choices[0].Message.Content = content
		if len(toolCalls) > 0 {
			resp.Choices[0].Message.ToolCalls = ToOpenAIToolCalls(toolCalls)
			resp.Choices[0].FinishReason = "tool_calls"
		}
	}

	return resp, nil
}

// StreamChat 处理流式聊天请求
// supportsNative: 模型是否支持原生 function call
func (p *Proxy) StreamChat(ctx context.Context, req *types.ChatRequest, supportsNative bool) (<-chan types.StreamChunk, error) {
	// 如果模型支持原生 function call 或没有工具，直接转发
	if supportsNative || len(req.Tools) == 0 {
		return p.client.StreamChat(ctx, req)
	}

	// 不支持原生 function call，需要模拟
	return p.handleSimulatedToolsStream(ctx, req)
}

// handleSimulatedToolsStream 处理流式模拟工具调用
func (p *Proxy) handleSimulatedToolsStream(ctx context.Context, req *types.ChatRequest) (<-chan types.StreamChunk, error) {
	// 1. 使用 renderer 将工具定义转换为提示词
	modifiedReq := p.renderer.RenderTools(req)

	// 2. 调用模型获取流式响应
	sourceCh, err := p.client.StreamChat(ctx, modifiedReq)
	if err != nil {
		return nil, fmt.Errorf("call model: %w", err)
	}

	// 3. 创建输出 channel 并启动解析 goroutine
	outputCh := make(chan types.StreamChunk, 10)

	go func() {
		defer close(outputCh)

		streamParser := parser.NewStreamParser("xml") // 使用 XML 流式解析器
		var fullContent string

		for chunk := range sourceCh {
			if chunk.Error != nil {
				outputCh <- chunk
				continue
			}

			// 处理每个 choice
			for i := range chunk.Choices {
				deltaContent := chunk.Choices[i].Delta.Content

				// 使用流式解析器解析
				toolCalls, content, _ := streamParser.Add(deltaContent)

				// 更新 chunk
				chunk.Choices[i].Delta.Content = content
				fullContent += deltaContent

				// 如果解析到工具调用，添加到 delta
				if len(toolCalls) > 0 {
					chunk.Choices[i].Delta.ToolCalls = ToOpenAIToolCalls(toolCalls)
					chunk.Choices[i].FinishReason = "tool_calls"
				}
			}

			// 发送修改后的 chunk
			select {
			case outputCh <- chunk:
			case <-ctx.Done():
				return
			}
		}

		// 刷新解析器，获取剩余内容
		allToolCalls, remaining := streamParser.Flush()
		if remaining != "" || len(allToolCalls) > 0 {
			finalChunk := types.StreamChunk{
				ID:      "chatcmpl-" + randomString(10),
				Object:  "chat.completion.chunk",
				Created: time.Now().Unix(),
				Model:   req.Model,
				Choices: []types.ChunkChoice{
					{
						Index: 0,
						Delta: types.Message{
							Content:   remaining,
							ToolCalls: ToOpenAIToolCalls(allToolCalls),
						},
						FinishReason: func() string {
							if len(allToolCalls) > 0 {
								return "tool_calls"
							}
							return "stop"
						}(),
					},
				},
			}
			select {
			case outputCh <- finalChunk:
			case <-ctx.Done():
			}
		}
	}()

	return outputCh, nil
}
