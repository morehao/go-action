package main

import (
	"fmt"
)

func main() {
	fmt.Println(movingCount(2, 3, 1))
}

func movingCount(m int, n int, k int) (ans int) {
	visited := &[100][100]bool{}
	return dfs(0, 0, m-1, n-1, k, visited)
}

func sum(x, y int) int {
	s := 0
	for x != 0 {
		// 通过取余获取个位数的值
		s += x % 10
		// 通过除法去除个位数
		x = x / 10
	}
	for y != 0 {
		s += y % 10
		y = y / 10
	}
	return s
}
func dfs(i, j, m, n, k int, visited *[100][100]bool) (ans int) {
	// 超出边界||已访问过
	if i > m || j > n || sum(i, j) > k || visited[i][j] {
		return 0
	}
	// 当前左边可达
	visited[i][j] = true
	// 坐标继续向右（j+1）和向下（i+1）
	return 1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited)
}
