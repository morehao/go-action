package main

import "fmt"

func main() {
	s := "你好"
	for _, v := range s {
		fmt.Printf("character %c,unicode %U, utf-8 %v, \n", v, v, v)
	}
}
