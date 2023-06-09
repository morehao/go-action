package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go processor(ch, i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

func processor(ch chan struct{}, i int) {
	fmt.Println(i)
	ch <- struct{}{}
}
