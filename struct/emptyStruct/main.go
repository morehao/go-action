package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		a := &struct{}{}
		fmt.Println(&a)
	}
}
