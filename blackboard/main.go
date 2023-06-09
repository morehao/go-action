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

func validateStackSequences(pushed []int, popped []int) bool {
	var (
		stack []int
		i     = 0
	)
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
