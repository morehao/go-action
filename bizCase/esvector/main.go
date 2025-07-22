package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("putVector", PutVector)
	r.GET("getVector", GetVector)

	r.Run(":8888")
}
