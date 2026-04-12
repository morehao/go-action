package main

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizcase/einodeer/config"
	"github.com/morehao/go-action/bizcase/einodeer/handler"
	"github.com/morehao/go-action/bizcase/einodeer/infra"
	"github.com/morehao/golib/conf"
	"github.com/morehao/golib/glog"
)

func main() {
	r := gin.Default()

	_, workDir, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(workDir)
	conf.SetAppRootDir(rootDir)
	config.LoadDeerConfig()

	infra.InitModel()
	// 初始化工具系统
	if err := infra.InitTools(); err != nil {
		glog.Errorf(context.Background(), "Failed to initialize tools: %v", err)
		panic(err)
	}

	r.POST("/api/chat/stream", handler.ChatStream)
	r.Run(":8080")
}
