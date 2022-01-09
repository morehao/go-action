package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // 设置计数器，即goroutine的个数
	go func() {
		//	Do some work ...
		time.Sleep(1 * time.Second)
		fmt.Print("Goroutine 1 finished!\n")
		wg.Done() // goroutine结束后将计数器减1
	}()
	go func() {
		//	Do some work ...
		time.Sleep(2 * time.Second)
		fmt.Print("Goroutine 2 finished!\n")
		wg.Done() // goroutine结束后将计数器减1
	}()
	wg.Wait() // 主goroutine阻塞，等待计数器变为0
	fmt.Print("All goroutine finished!")
}
