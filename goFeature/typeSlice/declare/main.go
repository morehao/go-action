package main

import "fmt"

func main() {
	// 	1、通过字面量创建切片
	s1 := []int{}
	s1 = append(s1, 1)
	fmt.Println(s1)
	// 	2、通过内置的make函数创建切片
	s2 := make([]int, 0)
	s2 = append(s2, 1)
	fmt.Println(s2)
}
