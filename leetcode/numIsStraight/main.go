package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 1, 2, 5}
	fmt.Println(isStraight(nums))
	fmt.Println(isStraight2(nums))
}

func isStraight(nums []int) bool {
	max, min := 0, 14
	m := make(map[int]bool)
	for i := range nums {
		if m[nums[i]] {
			return false
		}
		if nums[i] != 0 && nums[i] < min {
			m[nums[i]] = true
			min = nums[i]
		}
		if nums[i] != 0 && nums[i] > max {
			m[nums[i]] = true
			max = nums[i]
		}
	}
	// 刨除大小王，如果max-min<5，即可组成顺子
	return max-min < 5
}

func isStraight2(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	minIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			minIndex++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[minIndex] < 5
}
