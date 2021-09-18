package concurrent

import (
	"fmt"
	"time"
)

func MasterTaskQuit() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			// 主协程退出了，其他任务仍然执行
			break
		}
	}
}
