package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizcase/storagecase/config"
	"github.com/morehao/go-action/bizcase/storagecase/handler"
	"github.com/morehao/go-action/bizcase/storagecase/pkg/s3"
	"github.com/morehao/golib/glog"
	"gopkg.in/yaml.v3"
)

func main() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("storagecase: runtime.Caller failed")
	}
	rootDir := filepath.Dir(file)

	cfg, err := loadConfig(filepath.Join(rootDir, "config", "config.yaml"))
	if err != nil {
		glog.Errorf(context.Background(), "load config failed: %v", err)
		panic(err)
	}

	if err := s3.Init(cfg.Driver, cfg.Config); err != nil {
		glog.Errorf(context.Background(), "init storage failed: %v", err)
		panic(err)
	}

	h := handler.NewStorageHandler()

	r := gin.Default()

	api := r.Group("/api/storage")
	{
		api.GET("/health", h.Health)

		base := api.Group("/base")
		{
			base.POST("/put", h.PutObject)
			base.POST("/get", h.GetObject)
			base.POST("/delete", h.DeleteObject)
			base.POST("/delete-batch", h.DeleteObjects)
			base.POST("/list", h.ListObjects)
		}

		multipart := api.Group("/multipart")
		{
			multipart.POST("/create", h.CreateMultipartUpload)
			multipart.POST("/upload-part", h.UploadPart)
			multipart.POST("/complete", h.CompleteMultipartUpload)
			multipart.POST("/abort", h.AbortMultipartUpload)
		}

		ext := api.Group("/ext")
		{
			ext.POST("/head", h.HeadObject)
			ext.POST("/copy", h.CopyObject)
			ext.POST("/presign-get", h.PresignGetObject)
			ext.POST("/presign-put", h.PresignPutObject)
		}
	}

	glog.Infof(context.Background(), "storagecase starting on %s, driver=%s", cfg.ServerPort, cfg.Driver)
	if err := r.Run(cfg.ServerPort); err != nil {
		glog.Errorf(context.Background(), "server run failed: %v", err)
		panic(err)
	}
}

func loadConfig(path string) (*config.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg config.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	cfg.Defaults()
	return &cfg, nil
}
