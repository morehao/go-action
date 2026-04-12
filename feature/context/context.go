package context

import (
	"context"
	"fmt"
	"time"
)

// WithCancel WithCancel示例代码
func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go handle(ctx, 1*time.Second)
	time.Sleep(2 * time.Second)
}

// WithDeadline WithDeadline示例代码
func WithDeadline() {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("main err: ", ctx.Err())
	case <-time.After(500 * time.Millisecond):
		fmt.Println("main timeout")
	}
}

// WithTimeout WithTimeout示例代码
func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Println("start")

	select {
	case <-ctx.Done():
		fmt.Println("main err: ", ctx.Err())
	case <-time.After(500 * time.Millisecond):
		fmt.Println("main timeout")
	}
}

// WithValue WithValue示例代码
func WithValue() {
	ctx := context.WithValue(context.Background(), "key", "value")
	fmt.Println("v: ", ctx.Value("key"))
}

func handle(ctx context.Context, duration time.Duration) {
	go handle2(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("handle process request with", duration)
	}
}

func handle2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("handle2", ctx.Err())
	}
}
