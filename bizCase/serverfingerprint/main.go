package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	_ "github.com/morehao/golib/glog/driver/zap"
)

func main() {
	if err := glog.InitLogger(&glog.LogConfig{
		Service: "serverfingerprint",
		Writers: []glog.WriterConfig{{Type: glog.WriterConsole}},
	}); err != nil {
		panic(err)
	}
	defer glog.Close()

	r := gin.Default()

	r.GET("/health", HealthHandler)
	r.GET("/getSystemInfo", GetSystemInfoHandler)
	r.GET("/getFingerprint", GetFingerprint)

	r.Run(":8080")
}
