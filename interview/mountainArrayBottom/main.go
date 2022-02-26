package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1, 2, 3}
	fmt.Println(bottomIndexInMountainArray(nums))
}

func bottomIndexInMountainArray(nums []int) int {
	for i := 1; ; i++ {
		if nums[i] < nums[i+1] {
			return nums[i]
		}
	}
}
