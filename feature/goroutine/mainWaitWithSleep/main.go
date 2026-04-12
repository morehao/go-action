package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go processor(i)
	}
	time.Sleep(time.Second)
}

func processor(i int) {
	fmt.Println(i)
}
