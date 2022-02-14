package main

import (
	"fmt"
	"time"
)

func main() {
	tickerDemo()
}

func tickerDemo() {
	// 周期性地打印日志
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		fmt.Println("tickerDemo")
	}
}
