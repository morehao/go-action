package main

import (
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizCase/einodeer/config"
	"github.com/morehao/go-action/bizCase/einodeer/handler"
	"github.com/morehao/golib/conf"
)

func main() {
	r := gin.Default()

	_, workDir, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(workDir)
	conf.SetAppRootDir(rootDir)
	config.LoadDeerConfig()
	r.POST("/api/chat/stream", handler.ChatStreamEino)
	r.Run(":8888")
}
