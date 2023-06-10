package main

import (
	"fmt"
	"time"
)

// 输出main1 new1 new2 main2，主goroutine退出后其他goroutine也停止
func main() {
	// 合起来写
	go processor()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			// 主协程退出了，其他任务也退出了
			break
		}
	}
}

func processor() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(time.Second)
	}
}
