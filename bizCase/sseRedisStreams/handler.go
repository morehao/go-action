package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Chat(ctx *gin.Context) {
	hash := ctx.Query("hash")
	if hash == "" {
		ctx.JSON(400, gin.H{"error": "missing hash"})
		return
	}

	streamKey := "chat_stream_" + hash
	offsetKey := "offset_" + hash

	// 获取 offset
	lastID, _ := rdb.Get(ctx, offsetKey).Result()
	if lastID == "" {
		lastID = "0"
	}

	// 创建生成数据的通道
	ch := make(chan string, 100)

	// 查看是否已有数据
	length, _ := rdb.XLen(ctx, streamKey).Result()

	// 如果没有数据，启动写入任务（写 redis + 写 channel）
	if length == 0 {
		go MockModelStream(ctx.Copy(), ch)
	} else {
		close(ch) // 没有新任务生成，就关闭通道
	}

	// 设置 SSE 头
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Flush()

	clientDone := ctx.Request.Context().Done()

	for {
		select {
		case <-clientDone:
			log.Println("客户端断开，保存 offset", lastID)
			rdb.Set(ctx, offsetKey, lastID, time.Hour)
			return

		case data, ok := <-ch: // 优先从 MockModelStream 写入的 channel 读
			if !ok {
				ch = nil // 通道结束，关闭该 case
				continue
			}
			id := SaveToStream(ctx, streamKey, data)
			lastID = id

			fmt.Fprintf(ctx.Writer, "id: %s\ndata: %s\n\n", id, data)
			ctx.Writer.Flush()

			if data == "[DONE]" {
				rdb.Del(ctx, offsetKey)
				return
			}

		default:
			// 若通道已关闭，继续从 Redis Stream 读未发送的历史数据
			if ch == nil {
				msgs, _ := ReadFromStream(ctx, streamKey, lastID, 10, 3*time.Second)
				if len(msgs) == 0 {
					time.Sleep(300 * time.Millisecond)
					continue
				}
				for _, msg := range msgs {
					data := msg.Values["data"].(string)
					lastID = msg.ID
					fmt.Fprintf(ctx.Writer, "id: %s\ndata: %s\n\n", msg.ID, data)
					ctx.Writer.Flush()
					if data == "[DONE]" {
						rdb.Del(ctx, offsetKey)
						return
					}
				}
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
