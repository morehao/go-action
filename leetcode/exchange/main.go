package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}

func exchange(nums []int) []int {
	n := len(nums)
	res := make([]int, n, n)
	left, right := 0, n-1
	for _, v := range nums {
		if v%2 == 1 {
			res[left] = v
			left++
		} else {
			res[right] = v
			right--
		}
	}
	return res
}
