package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	process := func(i int, fn func()) {
		for {
			if count == i {
				fn()
				count++
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fn := func() {
				fmt.Println(i)
			}
			process(i, fn)
		}(i)
	}
	process(10, func() {})
}
