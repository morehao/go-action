package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const execNum = 3
	var wg sync.WaitGroup
	ch := make(chan struct{}, execNum)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			ch <- struct{}{}
			task(idx, ch)
		}(i)
	}
	wg.Wait()
}

func task(idx int, ch chan struct{}) {
	defer func() {
		<-ch
	}()
	fmt.Println("idx: ", idx)
	time.Sleep(time.Second * 2)
}
