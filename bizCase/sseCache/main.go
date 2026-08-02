package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	_ "github.com/morehao/golib/glog/driver/zap"
)

func main() {
	if err := glog.InitLogger(&glog.LogConfig{
		Service: "sseCache",
		Writers: []glog.WriterConfig{{Type: glog.WriterConsole}},
	}); err != nil {
		panic(err)
	}
	defer glog.Close()

	r := gin.Default()

	r.GET("/chat", Chat)
	r.GET("/stopChat", StopChat)
	r.GET("/message", GetMessage)

	r.Run(":8888")
}
