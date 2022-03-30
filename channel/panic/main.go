package main

import (
	"fmt"
	"time"
)

// 通道会发生panic的情况
func main() {
	go closeNilChan()
	go sendToClosedChan()
	go repeatClose()
	time.Sleep(1 * time.Second)
}

func closeNilChan() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("对nil通道进行关闭操作, err:", err)
		}
	}()
	var ch chan int
	// 对nil通道进行关闭操
	close(ch)
}

func sendToClosedChan() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("对已关闭的通道进行写操作, err:", err)
		}
	}()
	ch := make(chan int, 3)
	ch <- 2
	close(ch)
	// 对已关闭的通道进行发送操作
	ch <- 3
}

func repeatClose() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("对已关闭的通道进行再次关闭操作, err:", err)
		}
	}()
	ch := make(chan int, 3)
	ch <- 2
	close(ch)
	// 对已关闭的通道进行再次关闭操作
	close(ch)
}
