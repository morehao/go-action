package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	demoFn1()
	deadlock()
}

func demoFn1() {
	ch1, ch2 := make(chan int), make(chan int)
	select {
	case <-ch1:
		// 如果从 ch1 信道成功接收数据，则执行该分支代码
	case ch2 <- 1:
		// 如果成功向 ch2 信道成功发送数据，则执行该分支代码
	default:
		// 如果上面都没有成功，则进入 default 分支处理流程
	}
}

// select语句只能用于信道的读写操作
func caseFn() {
	size := 10
	ch := make(chan int, size)
	for i := 0; i < size; i++ {
		ch <- 1
	}
	ch2 := make(chan int, 1)

	select {
	// 表达式必须是channel的接收操作或写入操作
	// case 3 == 3:
	// 	fmt.Println("equal")
	case v := <-ch:
		fmt.Print(v)
	case ch2 <- 10:
		fmt.Print("write")
	default:
		fmt.Println("none")
	}
}

// 超时用法
func timeOut() {
	ch := make(chan int)
	go func(c chan int) {
		// 修改时间后,再查看执行结果
		time.Sleep(time.Second * 1)
		ch <- 1
	}(ch)

	select {
	case v := <-ch:
		fmt.Print(v)
	case <-time.After(2 * time.Second): // 等待 2s
		fmt.Println("no case ok")
	}

	time.Sleep(time.Second * 10)
}

// 空select
func emptySelect() {
	select {}
}

// for中的select 引起的CPU过高的问题
func deadlock() {
	quit := make(chan bool)
	for i := 0; i != runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-quit:
					break
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 15)
	for i := 0; i != runtime.NumCPU(); i++ {
		quit <- true
	}
}
