package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 让多个goroutine按顺序运行
func main()  {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			// 通过count与i对比，确保fn()执行的顺序，此处需了解原子操作相关函数
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	// I的类型指定为uint32是因为trigger函数中的原子操作对应函数入参为uint32
	for i := uint32(0); i < 10; i++ {
		// go语句传参i是为了确保每个goroutine可以拿到唯一的值
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	time.Sleep(time.Second)
}


