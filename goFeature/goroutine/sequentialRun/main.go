package main

import "fmt"

// 多个goroutine按顺序执行
func main() {
	channels := make([]chan int, 10)
	for i := range channels {
		channels[i] = make(chan int)
		go Process(channels[i])
	}
	for i, ch := range channels {
		<-ch
		fmt.Printf("goroutine %d quit \n", i)
	}
}

func Process(ch chan int) {
	ch <- 1
}
