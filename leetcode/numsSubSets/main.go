package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) (ans [][]int) {
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		// 回溯，没有nums[cur]时
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
