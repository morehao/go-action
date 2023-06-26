package main

import "fmt"

func main() {
	fn()
}

func fn() {
	numberCh, letterCh, doneCh := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go func() {
		i := 1
		for range numberCh {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letterCh <- struct{}{}
		}
	}()
	go func() {
		i := 'A'
		for range letterCh {
			if i >= 'Z' {
				doneCh <- struct{}{}
			} else {
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				numberCh <- struct{}{}
			}
		}
	}()
	numberCh <- struct{}{}
	for range doneCh {
		return
	}

}
