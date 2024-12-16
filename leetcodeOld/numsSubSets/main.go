package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
func subsets(nums []int) [][]int {
	var set []int
	var res [][]int
	var fn func(x int)
	fn = func(i int) {
		if len(nums) == i {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[i])
		fn(i + 1)
		// 回溯，没有nums[cur]时
		set = set[:len(set)-1]
		fn(i + 1)
	}
	fn(0)
	return res
}
