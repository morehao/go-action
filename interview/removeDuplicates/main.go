package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slowP := 0
	for fastP := 0; fastP < len(nums); fastP++ {
		if nums[fastP] != nums[slowP] {
			slowP++
			nums[slowP] = nums[fastP]
		}
	}
	return slowP + 1
}
