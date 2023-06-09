package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}
	return -1
}
