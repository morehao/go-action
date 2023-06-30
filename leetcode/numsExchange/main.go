package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}

func exchange(nums []int) []int {
	size := len(nums)
	res := make([]int, size, size)
	left, right := 0, size-1
	for i := range nums {
		if nums[i]%2 == 1 {
			res[left] = nums[i]
			left++
		} else {
			res[right] = nums[i]
			right--
		}
	}
	return res
}
