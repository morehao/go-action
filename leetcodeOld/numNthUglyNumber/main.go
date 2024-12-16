package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

/*
1、状态定义：设动态规划列表dp,dp[i]代表第i+1个丑数
2、转移方程：丑数是乘2乘3乘5的结果之一，取最小值即可,即dp[i] = min(dp[p2], dp[p3], dp[p5])
3、初始状态：第一个丑数为dp[0]=1
4、返回值：dp[n-1]表示第n个丑数
*/
func nthUglyNumber(n int) int {
	p2, p3, p5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2 := dp[p2] * 2
		n3 := dp[p3] * 3
		n5 := dp[p5] * 5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			p2++
		}
		if dp[i] == n3 {
			p3++
		}
		if dp[i] == n5 {
			p5++
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
