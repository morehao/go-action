package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(grid))
}

/*
根据题目说明，易得某单元格只可能从上边单元格或左边单元格到达。
设f(i, j)为从棋盘左上角走至单元格(i,j)的礼物最大累计价值，易得到以下递推关系：
f(i,j)等于f(i,j-1)和f(i-1,j)中的较大值加上当前单元格礼物价值grid(i，j）。
公式：f(i,j)=max(f(i,j-1),f(i-1,j))+grid(i,j)

	1、状态定义：dp[i,j]为从棋盘左上角走至单元格(i,j)的礼物最大累计价值
	2、转移方程：
		当i=0 且j=0 时，为起始元素，即d(i,j) = grid[i][j];
		当i=0 且j!=0 时，为矩阵第一行元素，只可从左边到达，即d(i,j) = d(i, j-1) + grid[i][j]；
		当i!=0 且j=0 时，为矩阵第一列元素，只可从上边到达，即d(i,j) = d(i-1, j) + grid[i][j]；
		当i!=0 且j!=0 时，可从左边或上边到达，即d(i,j) = max(d(i, j-1), d(i-1, j)) + grid[i][j]；
	3、初始状态：d[0,0] = grid[0][0]
	4、确定返回值：m和n分别为行高和列宽，遍历至右下角，所以返回为d[m-1][n-1]
*/
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[m-1][n-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
