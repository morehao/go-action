package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			// go的程序完全不会等待go函数执行完毕，所以不会有任何输出
			fmt.Println(i)
		}()
	}
}
