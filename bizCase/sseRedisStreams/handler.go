package main

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
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
		go MockModelStream(ctx.Copy(), hash)
	} else {
		close(ch) // 没有新任务生成，就关闭通道
	}

	// 设置 SSE 头
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Flush()

	clientGone := ctx.Stream(func(w io.Writer) bool {
		// 阻塞等待 Redis 有新消息推送
		msgs, err := ReadFromStream(ctx, streamKey, lastID, 1000, 5*time.Second)
		if err != nil {
			glog.Infof(ctx, "[Chat] read fail, error: %s", err)
			return false
		}
		if len(msgs) == 0 {
			return false
		}
		for _, msg := range msgs {
			data, ok := msg.Values["data"].(string)
			if !ok {
				continue
			}
			lastID = msg.ID

			glog.Infof(ctx, "[Chat] send data: %s", data)
			dataMsg := fmt.Sprintf(`{"id": %s, "data": "%s"}`, msg.ID, data)
			ctx.SSEvent("message", dataMsg)

			if data == "[DONE]" {
				rdb.Del(ctx, offsetKey)
				return false
			}
		}
		return false
	})

	if clientGone {
		glog.Infof(ctx, "[Chat] Client disconnected during streaming")
		rdb.Set(ctx, offsetKey, lastID, time.Hour)
	} else {
		glog.Infof(ctx, "[Chat] Stream completed normally")
	}

	return
}
