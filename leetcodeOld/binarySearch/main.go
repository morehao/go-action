package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(arr, 9))
}

// 二分法查找
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
