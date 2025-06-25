package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ygpkg/yg-go/apis/runtime/middleware"
)

func main() {
	r := gin.Default()

	r.GET("/chat", middleware.SSEHeader(), Chat)
	r.GET("/stopChat", middleware.SSEHeader(), StopChat)
	r.GET("/message", GetMessage)

	r.Run(":8888")
}
