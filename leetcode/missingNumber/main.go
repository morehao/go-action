package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(missingNumber(nums))
}

// 二分法
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

// 利用map，查找下标与值不相等的数字
func missingNumberWithMap(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}
	for i := 0; ; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
}

// 直接遍历
func missingNumberWithRange(nums []int) int {
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	// 长度为n-1，则遍历时下标最大为n-2，如果遍历过程中i!=v,则说明确实
	return len(nums)
}
