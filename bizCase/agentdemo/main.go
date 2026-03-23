package main

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizCase/agentdemo/agent"
	"github.com/morehao/go-action/bizCase/agentdemo/config"
	"github.com/morehao/go-action/bizCase/agentdemo/handler"
)

func main() {
	ctx := context.Background()

	// Resolve the directory that contains this source file so that relative
	// paths (e.g. config.yaml) always work regardless of the working directory.
	_, file, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(file)

	config.Load(filepath.Join(rootDir, "config", "config.yaml"))

	if err := agent.Init(ctx); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/chat", handler.Chat)
	r.POST("/chat/stream", handler.ChatStream)

	if err := r.Run(config.Cfg.Server.Port); err != nil {
		panic(err)
	}
}
