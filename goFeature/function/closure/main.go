package main

import "fmt"

func main() {
	add, sub := closure(10)
	// base初始值为10
	fmt.Println(add(1), sub(2))
	// 此时base是9
	fmt.Println(add(3), sub(4))
}

// 闭包
func closure(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}
