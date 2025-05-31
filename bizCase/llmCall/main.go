package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/chat", Chat)
	r.GET("/streamChat", StreamChat)
	r.GET("/functionCall", FunctionCall)
	r.Run(":8080")
}
