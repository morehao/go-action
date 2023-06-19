package main

import "fmt"

func main() {
	fmt.Println(sumNums(9))
}

func sumNums(n int) int {
	var sum func(x int)
	res := 0
	sum = func(x int) {
		res += x
		if x == 0 {
			return
		}
		sum(x - 1)
	}
	sum(n)
	return res
}
