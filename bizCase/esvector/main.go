package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	_ "github.com/morehao/golib/glog/driver/zap"
)

func main() {
	if err := glog.InitLogger(&glog.LogConfig{
		Service: "esvector",
		Writers: []glog.WriterConfig{{Type: glog.WriterConsole}},
	}); err != nil {
		panic(err)
	}
	defer glog.Close()

	r := gin.Default()

	r.GET("insertData", InsertData)
	r.GET("clearData", ClearData)
	r.POST("searchData", SearchData)

	r.Run(":8888")
}
