package main

import "fmt"

func main() {
	nums := []int{1, 9, 9}
	fmt.Println(plusOne(nums))
}

func plusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = (nums[i] + 1) % 10
		if nums[i] != 0 {
			return nums
		}
	}
	newNums := make([]int, len(nums)+1)
	newNums[0] = 1
	return newNums
}
