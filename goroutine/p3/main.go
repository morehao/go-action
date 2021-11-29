package main

import (
	"fmt"
	//"time"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1,简单粗暴，设置睡眠时间
	//time.Sleep(time.Millisecond * 500)

	// 办法2，缓冲通道缓冲区没有数据时，从通道读取数据会阻塞，直至有数据写入通道，即goroutine写数据到通道时主goroutine才能解除阻塞继续执行
	for j := 0; j < num; j++ {
		<-sign
	}
}
