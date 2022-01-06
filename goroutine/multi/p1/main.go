package main

import (
	"fmt"
	"sync"
	"time"
)

/*方案一：使用sync加chan控制并发数量*/
var wg = sync.WaitGroup{}

func main() {
	userCount := 10
	ch := make(chan bool, 4)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(ch, i)
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done()

	ch <- true
	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}
