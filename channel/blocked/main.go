package main

// 通道会发生阻塞的情况
func main() {
	sendToNilChan()
	// readFromNilChan()
	// operateUnbufferedChan()
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

// 非缓冲队列的读写操作会导致永久阻塞
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
