package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("i:", i)
		}()
	}
}
