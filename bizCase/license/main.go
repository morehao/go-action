package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// r.GET("/health", HealthHandler)
	// r.GET("/getSystemInfo", GetSystemInfoHandler)
	// r.GET("/getFingerprint", GetFingerprint)

	r.Run(":8080")
}
