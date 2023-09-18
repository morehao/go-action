package main

import (
	"log"
	"net/http"

	"go-homework/router"
)

func main() {
	HttpServer()
}

func HttpServer() {
	router.Http()
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
