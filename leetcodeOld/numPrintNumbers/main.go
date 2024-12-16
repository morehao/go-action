package main

import (
	"fmt"
)

func main() {
	fmt.Println(printNumbers(2))
}

func printNumbers(n int) []int {
	var res []int
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	for i := 1; i < max; i++ {
		res = append(res, i)
	}
	return res
}
