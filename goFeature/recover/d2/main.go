package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()
	panic("demo")
	fmt.Println("continue running")
}
