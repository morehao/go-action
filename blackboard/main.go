package main

import "fmt"

func main() {
	fmt.Println(printNumbers(2))
}

func printNumbers(n int) []int {
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	var res []int
	for i := 0; i < max; i++ {
		res = append(res, i)
	}
	return res
}
