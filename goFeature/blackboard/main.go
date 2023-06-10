package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
	}()
	fmt.Println("running")
	cancel()
	select {
	case <-ctx.Done():
		fmt.Println("context cancelled")
	case <-time.After(4 * time.Second):
		fmt.Println("timeout")
	}
}

func search(nums []int, target int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if v == target {
			m[target]++
		}
	}
	return m[target]
}
