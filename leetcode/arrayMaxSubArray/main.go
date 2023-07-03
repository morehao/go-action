package main

import "fmt"

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}

/*
动态规划
1、状态定义：dp[i]代表以nums[i]结尾的连续数组的最大和
2、转移方程：如果dp[i-1]<0，那么dp[i-1]对于dp[i]是负的贡献，还不如nums[i]大，所以公式如下：
		当dp[i-1]>0时：执行dp[i]=dp[i-1]+nums[i];
		当dp[i-1]≤0时：执行dp[i]=nums[i];
3、初始状态：dp[0]=nums[0]
*/
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			// 直接修改原值，节省空间
			nums[i] += nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
