package main

import "fmt"

func main() {
	res := findTargetIndex([]int{1, 3, 5, 8, 7}, 8)
	fmt.Printf("result:%v", res)
}

// 找出两数之和等于目标值的下标
func findTargetIndex(arr []int, target int) []int {
	indexMap := make(map[int]int, 0)
	for k, v := range arr {
		diff := target - v
		_, ok := indexMap[diff]
		if ok {
			return []int{k, indexMap[diff]}
		}
		indexMap[v] = k
	}
	return []int{}
}

// 双指针，找出具体元素
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	tmp := 0
	for i < j {
		tmp = nums[i] + nums[j]
		if tmp == target {
			return []int{nums[i], nums[j]}
		} else if tmp > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
