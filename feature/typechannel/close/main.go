package main

import "fmt"

func main() {
	// isClose()
	rangeCloseCh()
}

func isClose() {
	ch := make(chan int, 4)
	ch <- 1
	close(ch)
	if v, ok := <-ch; ok {
		fmt.Println(v)
	}
}

func rangeCloseCh() {
	c := make(chan int, 10)
	c <- 2
	c <- 3
	// close(c)
	for i := range c {
		fmt.Println(i)
	}
}
