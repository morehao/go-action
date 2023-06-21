package main

import "fmt"

func main() {
	n := 5
	fmt.Println(numWays(n))
}

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	const mod = 1000000007
	pre, curr := 1, 1
	for i := 2; i <= n; i++ {
		next := pre + curr
		pre = curr
		curr = next % mod
	}
	return curr
}
