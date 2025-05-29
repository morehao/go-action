package main

import (
	"fmt"
	"time"
)

// 生产者每 1 秒生产一条消息，消费者每 2 秒消费一条消息（比生产者慢），也要考虑如何加快消费
func main() {
	ch := make(chan string, 100)
	// 每一个 goroutine 相当于一个消费者，需要增加消费者数量，就增加 go 调用
	// 消费者 1
	go func() {
		for {
			time.Sleep(time.Second * 2)
			select {
			case a := <-ch:
				fmt.Println("data: ", a)
			default:
				fmt.Println("queue is empty")

			}
		}
	}()
	// 消费者 2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			select {
			case a := <-ch:
				fmt.Println("data: ", a)
			default:
				fmt.Println("queue is empty")

			}
		}
	}()
	// 生产者
	for {
		time.Sleep(time.Second)
		select {
		case ch <- time.Now().String():
		default:
			fmt.Println("queue is full")
		}
	}
}
