package main

import "fmt"

func main() {
	fmt.Printf("%d\n", fib(5))
	fmt.Println("-----------分割符-------")
	fmt.Printf("%d\n", fib2(5))
}

// 递归-斐波那切数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	// 需要取模时
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	pre, current := 0, 1
	for i := 2; i <= n; i++ {
		next := pre + current
		pre = current
		current = next % mod
	}
	return current
}
