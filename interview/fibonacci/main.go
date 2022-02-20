package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fibonacci(i))
	}
}

// 递归-斐波那切数列
func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
