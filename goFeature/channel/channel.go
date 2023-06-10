package channel

import "fmt"

func InitChannel() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	// 先进先出，输出2
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
}

func NoReceiverChannel() {
	// 未声明容量，是无缓冲的通道，无缓冲的通道只有在有接收者的时候才能发送值
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func WithReceiverChannel() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

func Close() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

func Panic() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	close(ch1)
	// 对已关闭的通道进行发送操作
	ch1 <- 3
	// 对已关闭的通道进行再次关闭操作
	close(ch1)
}

// TODO:待确认如何算是优雅
func GetValGoodExample() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	close(ch1)
	val, ok := <-ch1
	// 判断通道是否关闭，未关闭时再写入值
	if ok {
		ch1 <- 2
		fmt.Println(val)
	}
}
