package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i)
		fmt.Println(nums[i])
	}
}
func getLeastNumbers(arr []int, k int) []int {
	partition(arr, 0, len(arr)-1)
	return arr[len(arr)-k:]
}

func partition(nums []int, start, end int) {
	i, j := start, end
	midValue := nums[(i+j)/2]
	for i <= j {
		for nums[i] < midValue {
			i++
		}
		for nums[j] > midValue {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if i < end {
		partition(nums, i, end)
	}
	if j > start {
		partition(nums, start, j)
	}
	return
}
