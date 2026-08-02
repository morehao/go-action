package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	_ "github.com/morehao/golib/glog/driver/zap"
)

func main() {
	if err := glog.InitLogger(&glog.LogConfig{
		Service: "sseredisstreams",
		Writers: []glog.WriterConfig{{Type: glog.WriterConsole}},
	}); err != nil {
		panic(err)
	}
	defer glog.Close()

	r := gin.Default()

	r.GET("/chat", Chat)

	r.Run(":8080")
}
