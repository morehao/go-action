package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
