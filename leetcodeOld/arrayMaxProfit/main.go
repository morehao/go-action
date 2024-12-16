package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	fmt.Println(violentMaxProfit(prices))
}

/*
动态规划
1.状态定义：设动态规划列表dp,dp[i]代表以prices[i]为结尾的子数组的最大利润（以下简称为前i日的最大利润)
2、转移方程：由于题目限定“买卖该股票一次”，因此前日最大利润dp⑦等于前。一1日最大利润 dp[i一1和第i日卖出的最大利润中的最大值。
前i日最大利润=max(前(i一1)日最大利润，第日价格一前日最低价格) dp[i=max(dp[i-1],prices[i-min(prices0:i))
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var (
		profit   = 0
		minPrice = prices[0]
	)
	for _, price := range prices {
		minPrice = min(minPrice, price)
		profit = max(profit, price-minPrice)
	}
	return profit
}

// 暴力解法，找出最大间距
func violentMaxProfit(prices []int) int {
	var res int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
