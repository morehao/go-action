/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



 示例 1：


输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

 示例 2：


输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶




 提示：


 1 <= n <= 45


 Related Topics 记忆化搜索 数学 动态规划 👍 3741 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
// 动态规划
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	// f(x)=f(x−1)+f(x−2)
	pre, curr := 1, 1
	for i := 2; i <= n; i++ {
		next := pre + curr
		pre = curr
		curr = next
	}
	return curr
}

// leetcode submit region end(Prohibit modification and deletion)
