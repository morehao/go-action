/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

 问总共有多少条不同的路径？



 示例 1：


输入：m = 3, n = 7
输出：28

 示例 2：


输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下


 示例 3：


输入：m = 7, n = 3
输出：28


 示例 4：


输入：m = 3, n = 3
输出：6



 提示：


 1 <= m, n <= 100
 题目数据保证答案小于等于 2 * 10⁹


 Related Topics 数学 动态规划 组合数学 👍 2196 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	// 创建一个大小为 n 的一维数组 dp，用于存储当前行的路径数
	dp := make([]int, n)

	// 初始化 dp 数组，第一行的每个位置只有一种路径（一直向右走）
	for i := range dp {
		dp[i] = 1
	}

	// 动态规划填充 dp 数组
	for i := 1; i < m; i++ { // 从第二行开始
		for j := 1; j < n; j++ { // 从第二列开始
			// 状态转移方程：dp[j] = dp[j] + dp[j-1]
			// dp[j] 表示从上方 (i-1, j) 到达 (i, j) 的路径数（即上一行的 dp[j]）
			// dp[j-1] 表示从左方 (i, j-1) 到达 (i, j) 的路径数（即当前行的 dp[j-1]）
			dp[j] += dp[j-1]
		}
	}

	// 返回从起点 (0, 0) 到终点 (m-1, n-1) 的路径数
	return dp[n-1]
}

// leetcode submit region end(Prohibit modification and deletion)
