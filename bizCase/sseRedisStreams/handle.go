package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// 处理 SSE 连接
func handleSSE(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	// 设置 SSE 相关的 header
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")

	// 获取用户上次的消息 ID
	lastMsgID, err := rdb.Get(ctx, fmt.Sprintf("user:%s:last_msg_id", userID)).Result()
	if errors.Is(err, redis.Nil) {
		lastMsgID = "0" // 如果是新用户，从头开始读取
	} else if err != nil {
		log.Printf("Error getting last message ID: %v", err)
		lastMsgID = "0"
	}

	// 创建一个通道用于检测客户端断开连接
	cleanup := ctx.Writer.CloseNotify()

	for {
		select {
		case <-cleanup:
			// 客户端断开连接，保存最后读取的消息 ID
			if lastMsgID != "0" {
				err := rdb.Set(ctx, fmt.Sprintf("user:%s:last_msg_id", userID), lastMsgID, 24*time.Hour).Err()
				if err != nil {
					log.Printf("Error saving last message ID: %v", err)
				}
			}
			return

		default:
			// 从 Redis Stream 中读取消息
			streams, err := rdb.XRead(ctx, &redis.XReadArgs{
				Streams: []string{"messages", lastMsgID},
				Count:   1,
				Block:   0,
			}).Result()

			if err != nil {
				log.Printf("Error reading from stream: %v", err)
				continue
			}

			// 处理消息
			for _, stream := range streams {
				for _, message := range stream.Messages {
					lastMsgID = message.ID
					content := message.Values["content"].(string)

					// 发送消息到客户端
					ctx.SSEvent("message", gin.H{
						"id":      message.ID,
						"content": content,
					})
					ctx.Writer.Flush()
				}
			}
		}
	}
}

// 添加消息到 Stream
func addMessage(ctx *gin.Context) {
	content := ctx.PostForm("content")
	if content == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "content is required"})
		return
	}

	// 添加消息到 Redis Stream
	msgID, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "messages",
		Values: map[string]interface{}{
			"content": content,
		},
	}).Result()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message_id": msgID,
		"content":    content,
	})
}
