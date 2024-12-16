package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	max := 1
	var res int
	for i := range nums {
		if i == 0 {
			res = nums[i]
		}
		m[nums[i]]++
		if m[nums[i]] > max {
			max = m[nums[i]]
			res = nums[i]
		}
	}
	return res
}
