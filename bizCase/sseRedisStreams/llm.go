package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MockModelStream(ctx *gin.Context, ch chan string) {
	defer close(ch)

	for i := 0; i < 200; i++ {
		select {
		case <-ctx.Request.Context().Done():
			return
		default:
			data := fmt.Sprintf("data lind: %d", i)
			ch <- data
			time.Sleep(300 * time.Millisecond)
		}
	}
	ch <- "[DONE]"
}
