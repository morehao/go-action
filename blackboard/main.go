package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包
)

func main() {
	go func() {
		// 启动一个 HTTP 服务器，监听 localhost:6060 端口，暴露 pprof 相关的信息。
		log.Println(http.ListenAndServe("localhost:6060", nil)) // 启动 pprof 服务
	}()

	// 模拟代码中的字符串截取操作
	var str0 = "12345678901234567890"
	str1 := str0[:10]

	fmt.Println(str1)

	// 阻塞主程序，让它运行，直到用户手动终止
	select {}
}
