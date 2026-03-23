package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
)

type chatRequest struct {
	Message string `json:"message" binding:"required"`
}

// Chat handles POST /chat requests and streams the agent's response via SSE.
func Chat(ctx *gin.Context) {
	var req chatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "text/event-stream; charset=utf-8")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	iter := agentRunner.Run(ctx.Request.Context(),
		[]adk.Message{schema.UserMessage(req.Message)})

	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			writeSSEEvent(ctx, "error", event.Err.Error())
			break
		}
		if event.Output == nil || event.Output.MessageOutput == nil {
			continue
		}
		mv := event.Output.MessageOutput
		if mv.Role != schema.Assistant {
			continue
		}
		if mv.MessageStream != nil {
			for {
				chunk, err := mv.MessageStream.Recv()
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					writeSSEEvent(ctx, "error", fmt.Sprintf("stream error: %v", err))
					break
				}
				if chunk != nil && chunk.Content != "" {
					writeSSEEvent(ctx, "message", chunk.Content)
				}
			}
		} else if mv.Message != nil && mv.Message.Content != "" {
			writeSSEEvent(ctx, "message", mv.Message.Content)
		}
	}

	_, _ = fmt.Fprintf(ctx.Writer, "data: [DONE]\n\n")
	ctx.Writer.Flush()
}

func writeSSEEvent(ctx *gin.Context, event, data string) {
	_, _ = fmt.Fprintf(ctx.Writer, "event: %s\ndata: %s\n\n", event, data)
	ctx.Writer.Flush()
}

// Healthcheck handles GET /healthcheck.
func Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
