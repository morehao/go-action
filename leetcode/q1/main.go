package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	letter, number := make(chan struct{}), make(chan struct{})
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d", i)
				i++
				fmt.Printf("%d", i)
				i++
				letter <- struct{}{}
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- struct{}{}
			}
		}
	}(&wg)
	number <- struct{}{}
	wg.Wait()
}
