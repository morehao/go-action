package main

import (
	"fmt"
	"time"
)

func main() {
	// panic1()
	panic2()
}

// panic不能被跨协程的recover捕获
// panic1函数由新的协程调度，panic不能被主goroutine的recover捕获
func panic1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover_panic")
		}
	}()

	go func() {
		panic("goroutine2_panic")
	}()

	time.Sleep(1 * time.Second)
}

// 一个协程panic之后，是会导致所有的协程全部挂掉的，程序会整体退出
func panic2() {
	// 协程A
	go func() {
		for {
			fmt.Println("goroutine1_print")
		}
	}()

	// 协程B
	go func() {
		time.Sleep(1 * time.Second)
		panic("goroutine2_panic")
	}()

	time.Sleep(2 * time.Second)
}
