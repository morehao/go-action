package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MockModelStream(ctx *gin.Context, hash string) {
	streamKey := "chat_stream_" + hash
	for i := 0; i < 200; i++ {
		data := fmt.Sprintf("data lind %d", i)
		SaveToStream(ctx, streamKey, data)
		time.Sleep(300 * time.Millisecond)
	}
	SaveToStream(ctx, streamKey, "[DONE]")
}
