package router

import (
	"net/http"

	"go-homework/controller/ctrUser"
)

func Http() {
	http.HandleFunc("/healthz", ctrUser.Healthz) // 设置访问的路由
}
