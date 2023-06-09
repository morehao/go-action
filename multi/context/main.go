package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// CancelCtxFn()
	// DeadlineCtxFn()
	TimeoutCtxFn()
}

func CancelCtxFn() {
	// 使用context.WithCancel创建context时，如果没有父context，则需要传入background作为其父节点。
	ctx, cancel := context.WithCancel(context.Background())
	go handleRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Print("It's time to stop all sub goroutines!\n")
	cancel()
	// 只是为了测试子goroutine是否存在
	time.Sleep(5 * time.Second)
}

func DeadlineCtxFn() {
	// 使用context.WithDeadline创建context时，如果没有父context，则需要传入background作为其父节点。
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	go handleRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Print("It's time to stop all sub goroutines!\n")
	cancel()
	// 只是为了测试子goroutine是否存在
	time.Sleep(5 * time.Second)
}

func TimeoutCtxFn() {
	// 使用context.WithTimeout创建context时，如果没有父context，则需要传入background作为其父节点。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	go handleRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Print("It's time to stop all sub goroutines!\n")
	cancel()
	// 只是为了测试子goroutine是否存在
	time.Sleep(5 * time.Second)
}

func handleRequest(ctx context.Context) {
	go writeRedis(ctx)
	go writeDataBase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Print("handleRequest done.\n")
			return
		default:
			fmt.Print("handleRequest running.\n")
			time.Sleep(2 * time.Second)
		}
	}
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

func writeDataBase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("writeDataBase done.\n")
			return
		default:
			fmt.Print("writeDataBase running.\n")
			time.Sleep(2 * time.Second)
		}
	}
}
