package main

import (
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizCase/einodeer/config"
	"github.com/morehao/go-action/bizCase/einodeer/handler"
	"github.com/morehao/go-action/bizCase/einodeer/infra"
	"github.com/morehao/golib/conf"
	"github.com/morehao/golib/glog"
)

func main() {
	r := gin.Default()

	_, workDir, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(workDir)
	conf.SetAppRootDir(rootDir)
	config.LoadDeerConfig()
	
	// 初始化工具系统
	if err := infra.InitTools(); err != nil {
		glog.Errorf(nil, "Failed to initialize tools: %v", err)
		panic(err)
	}
	
	r.POST("/api/chat/stream", handler.ChatStreamEino)
	r.Run(":8888")
}
