package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/chat", Chat)

	r.Run(":8080")
}
