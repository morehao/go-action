/**
给定一个 m x n 的整数数组 obstacleGrid。一个机器人初始位于 左上角（即 obstacleGrid[0][0]）。机器人尝试移动到 右下角（即 obstacleGrid[m - 1][n -
 1]）。机器人每次只能向下或者向右移动一步。

 网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。

 返回机器人能够到达右下角的不同路径数量。

 测试用例保证答案小于等于 2 * 10⁹。



 示例 1：


输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右


 示例 2：


输入：obstacleGrid = [[0,1],[0,0]]
输出：1




 提示：


 m == obstacleGrid.length
 n == obstacleGrid[i].length
 1 <= m, n <= 100
 obstacleGrid[i][j] 为 0 或 1


 Related Topics 数组 动态规划 矩阵 👍 1388 👎 0

*/

package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 如果起点或终点是障碍物，则无法到达
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// 定义 dp 数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起点
	dp[0][0] = 1

	// 按列处理（如果有障碍物，后续全为 0）
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	// 按行处理（如果有障碍物，后续全为 0）
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	// 计算 dp 数组
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// leetcode submit region end(Prohibit modification and deletion)
