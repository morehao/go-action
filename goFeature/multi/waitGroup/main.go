package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	demo3()
}

// 简单使用sync.WaitGroup实现并发
func demo1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("goroutine 1 finished.")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("goroutine 2 finished.")
	}()
	wg.Wait()
	fmt.Println("all goroutine finished.")
}

// 将sync.WaitGroup和channel简单封装，可以支持并发数控制
func demo2() {
	multi := NewMulti(2)
	for i := 0; i < 10; i++ {
		go demo2Processor(multi, i)
	}
	multi.Wait()
}

func demo2Processor(m *Multi, i int) {
	defer m.Done()
	m.Add(1)
	fmt.Printf("goroutine %d finished. \n", i)
	time.Sleep(time.Second)
}

// multi封装Run方法
func demo3() {
	multi := NewMulti(2)
	for i := 0; i < 10; i++ {
		tempI := i
		multi.Run(func() {
			demo3Processor(tempI)
		})
	}
	multi.Wait()
}

func demo3Processor(i int) {
	fmt.Printf("goroutine %d finished. \n", i)
	time.Sleep(time.Second)
}
