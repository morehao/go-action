package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover err:%v", err)
		}
	}()
	panic("demo")
}
