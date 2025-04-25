package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用场景1，设定超时时间
	channel := make(chan string, 1)
	channel <- "a"
	waitChannel(channel)
	// 使用场景2，延迟执行某个方法
	delayFunction()

	// 重置定时器
	resetTimer()
}

func waitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timer.Stop()
		fmt.Println("waitChannel get data!")
		return true
	case <-timer.C: // 超时
		fmt.Println("waitChannel timeout!")
		return false
	}
}

func delayFunction() {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("delayFunction run!")
	}
}

func resetTimer() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("newTimer now:%v\n", time.Now())
	timer.Stop()
	timer.Reset(2 * time.Second)
	select {
	case <-timer.C:
		fmt.Printf("resetTimer now:%v\n", time.Now())
	}
}
