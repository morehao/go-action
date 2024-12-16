package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 8))
	fmt.Println(searchWithMap(nums, 8))
}

// 利用map进行查询
func searchWithMap(nums []int, target int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if v == target {
			m[target]++
		}
	}
	return m[target]
}

// 二分查找法，找到target出现的下标，找到target+1出现的下标，相减即可
func search(nums []int, target int) int {
	left := findLeft(nums, target)
	right := findLeft(nums, target+1)
	return right - left
}

func findLeft(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
