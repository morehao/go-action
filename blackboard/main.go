package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go writeRedis(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func writeRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("writeRedis done.\n")
			return
		default:
			fmt.Print("writeRedis running.\n")
			time.Sleep(2 * time.Second)
		}
	}
}

func waitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go fn1(&wg, i)
	}
	wg.Wait()
}

func fn1(wg *sync.WaitGroup, i int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("i:%d\n", i)
	wg.Done()
}
