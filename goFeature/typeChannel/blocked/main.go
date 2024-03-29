package main

import "fmt"

// 通道会发生阻塞的情况
func main() {
	// sendToNilChan()
	// readFromNilChan()
	// operateUnbufferedChan()
	rangeEmptyBufferedChan()
}

// 向nil通道中写入会导致永久阻塞
func sendToNilChan() {
	var ch chan int
	ch <- 1
}

// 从nil通道中读取会导致永久阻塞
func readFromNilChan() {
	var ch chan int
	<-ch
}

// 非缓冲通道的读写操作会导致永久阻塞
// 非缓冲通道发送时，在接受者准备好之前都是阻塞的；读操作在发送者准备好之前都是阻塞的。
// 如果需要满足发送者和接收者都准备好，可以开启新的goroutine让其中一个准备好
func operateUnbufferedChan() {
	ch := make(chan int)
	ch <- 1
	<-ch
	// 使用 goroutine 对非缓冲通道进行发送或读取操作，可以防止阻塞
	// go func() {
	// 	ch <- 1
	// }()
	// <-ch
}

// 对已空的缓冲通道继续读的操作
func rangeEmptyBufferedChan() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	<-ch
	<-ch
	<-ch
	fmt.Println("will block")
	// 	ch长度为3，读第四次时会发生阻塞
	<-ch
}
