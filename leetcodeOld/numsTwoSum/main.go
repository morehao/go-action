package main

import "fmt"

func main() {
	res := twoSum([]int{1, 3, 5, 8, 7}, 8)
	fmt.Printf("result:%v", res)
}

// 双指针，找出具体元素
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		tempSum := nums[i] + nums[j]
		if tempSum == target {
			return []int{nums[i], nums[j]}
		}
		if tempSum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

// 找出两数之和等于目标值的下标
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		diff := target - nums[i]
		if m[diff] > 0 {
			return []int{diff, m[diff]}
		}
		m[nums[i]] = diff
	}
	return nil
}
