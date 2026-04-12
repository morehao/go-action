// Package handler provides the Gin HTTP handlers for the agent demo.
package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizcase/agentdemo/agent"
)

// ChatRequest is the JSON body accepted by both /chat and /chat/stream.
type ChatRequest struct {
	Messages []MessageDTO `json:"messages" binding:"required,min=1"`
}

// MessageDTO is a single conversation turn.
type MessageDTO struct {
	Role    string `json:"role"    binding:"required"`
	Content string `json:"content" binding:"required"`
}

// ChatResponse is the JSON response returned by /chat.
type ChatResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

// Chat handles POST /chat – synchronous (non-streaming) inference.
func Chat(ctx *gin.Context) {
	req := new(ChatRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, ChatResponse{Code: 400, Message: err.Error()})
		return
	}

	msgs := toSchemaMessages(req.Messages)
	result, err := agent.GetAgent().Generate(ctx, msgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ChatResponse{Code: 500, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ChatResponse{Code: 0, Message: "success", Data: result.Content})
}

// ChatStream handles POST /chat/stream – SSE streaming inference.
func ChatStream(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/event-stream; charset=utf-8")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	req := new(ChatRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		writeSSE(ctx.Writer, "error", err.Error())
		return
	}

	msgs := toSchemaMessages(req.Messages)
	sr, err := agent.GetAgent().Stream(ctx, msgs)
	if err != nil {
		writeSSE(ctx.Writer, "error", err.Error())
		return
	}
	defer sr.Close()

	flusher, canFlush := ctx.Writer.(http.Flusher)
	for {
		msg, recvErr := sr.Recv()
		if recvErr != nil {
			if recvErr != io.EOF {
				writeSSE(ctx.Writer, "error", recvErr.Error())
			}
			break
		}
		if msg != nil && msg.Content != "" {
			writeSSE(ctx.Writer, "message", msg.Content)
			if canFlush {
				flusher.Flush()
			}
		}
	}
	writeSSE(ctx.Writer, "done", "[DONE]")
	if canFlush {
		flusher.Flush()
	}
}

// toSchemaMessages converts DTO messages into the schema types expected by Eino.
func toSchemaMessages(dtos []MessageDTO) []*schema.Message {
	msgs := make([]*schema.Message, 0, len(dtos))
	for _, d := range dtos {
		switch d.Role {
		case "system":
			msgs = append(msgs, schema.SystemMessage(d.Content))
		case "assistant":
			msgs = append(msgs, schema.AssistantMessage(d.Content, nil))
		default: // "user" or anything else
			msgs = append(msgs, schema.UserMessage(d.Content))
		}
	}
	return msgs
}

// writeSSE writes a single SSE frame to w.
func writeSSE(w http.ResponseWriter, event, data string) {
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, data)
}
