package main

import "fmt"

func main() {
	SliceExtend()
}

func SliceExtend() {
	var s []int
	s = append(s, 1)
	fmt.Println(s)
}
