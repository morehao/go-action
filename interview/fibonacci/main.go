package main

import "fmt"

func main() {
	fmt.Printf("%d\n", fibonacci(4))
	fmt.Println("-----------分割符-------")
	fmt.Printf("%d\n", fibonacci2(4))
}

// 递归-斐波那切数列
func fibonacci(i int) int {
	if i < 2 {
		return i
	}
	if i == 2 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func fibonacci2(i int) int {
	if i < 2 {
		return i
	}
	if i == 2 {
		return 1
	}
	pre, current, index := 1, 1, 3
	for index <= i {
		next := pre + current
		pre = current
		current = next
		index++
	}
	return current
}
