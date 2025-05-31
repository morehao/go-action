package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/newChat", NewChat)
	r.Run(":8080")
}
