package main

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	//	Do some work ...
	time.Sleep(time.Second)
	ch <- 1 // 在管道中写入一个元素，标识当前携程已经结束
}

func main() {
	channels := make([]chan int, 10) // 创建一个元素类型为channel长度为10的切片
	for i := range channels {
		channels[i] = make(chan int) // 在切片中放入一个channel
		go Process(channels[i])
	}
	for i, ch := range channels {
		<-ch
		fmt.Printf("Routine %d quit\n", i)
	}
}
