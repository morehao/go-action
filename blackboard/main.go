package main

import "fmt"

func main() {
	SliceExtend()
}

func SliceExtend() {
	var slice []int
	s1 := append(slice, 1, 2, 3)
	fmt.Println(cap(s1))
	s2 := append(s1, 4)
	fmt.Println(&s1[0] == &s2[0])
}
