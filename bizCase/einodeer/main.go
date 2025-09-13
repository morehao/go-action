package main

import (
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/conf"
)

func main() {
	r := gin.Default()

	_, workDir, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(workDir)
	conf.SetAppRootDir(rootDir)
	r.Run(":8888")
}
