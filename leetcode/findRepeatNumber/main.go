package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	var res int
	for _, v := range nums {
		if m[v] {
			res = v
			break
		}
		m[v] = true
	}
	return res
}
