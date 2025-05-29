package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// SSE 端点
	r.GET("/stream", handleSSE)

	// 添加消息的端点
	r.POST("/message", addMessage)

	// 启动服务器
	r.Run(":8080")
}
