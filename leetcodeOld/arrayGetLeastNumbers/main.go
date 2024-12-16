package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(getLeastNumbers(nums, 2))
	fmt.Println(getLeastNumbers2(nums, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
	quickSort(arr)
	return arr[:k]
}
func quickSort(nums []int) {
	partition(nums, 0, len(nums)-1)
}

func partition(nums []int, start, end int) {
	i, j := start, end
	midValue := nums[(start+end)/2]
	for i <= j {
		for nums[i] < midValue {
			i++
		}
		for nums[j] > midValue {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if i < end {
		partition(nums, i, end)
	}
	if j > start {
		partition(nums, start, j)
	}
}
