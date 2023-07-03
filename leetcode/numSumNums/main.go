package main

import "fmt"

func main() {
	fmt.Println(sumNums(9))
}

func sumNums(n int) int {
	ans := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		// 如果n>0为fasle，则后面不会进行递归调用
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
