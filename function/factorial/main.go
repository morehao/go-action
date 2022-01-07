package main

import "fmt"

func main() {
	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}

// 递归-阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
