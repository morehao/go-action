package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(getLeastNumbers(nums, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[:k]
}
