package main

import (
	"fmt"
	"log"
	"os"

	"excelagent/config"
	"excelagent/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if err := os.MkdirAll(cfg.Task.Workspace, 0755); err != nil {
		log.Fatalf("Failed to create workspace: %v", err)
	}
	if err := os.MkdirAll(cfg.Task.Workspace+"/input", 0755); err != nil {
		log.Fatalf("Failed to create input directory: %v", err)
	}

	r := gin.Default()

	taskHandler := handler.NewTaskHandler(cfg)
	taskHandler.RegisterRoutes(r)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
