package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	process := func(i int, fn func()) {
		for {
			if i == count {
				fn()
				count++
				break
			}
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
	time.Sleep(time.Second * 3)
}
