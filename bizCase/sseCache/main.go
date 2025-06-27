package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/chat", Chat)
	r.GET("/stopChat", StopChat)
	r.GET("/message", GetMessage)

	r.Run(":8888")
}
