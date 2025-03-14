/**
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。



 示例 1：


输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。


 示例 2：


输入：n = 0
输出：0


 示例 3：


输入：n = 1
输出：0




 提示：


 0 <= n <= 5 * 10⁶


 Related Topics 数组 数学 枚举 数论 👍 1211 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	set := make([]bool, n)
	for i := range set {
		set[i] = true
	}
	var count int
	for i := 2; i < n; i++ {
		if set[i] {
			count++
			// 如果i是质数，2*i则一定不是质数，把不是质数的标记出来，剩下的都是质数
			for j := 2 * i; j < n; j += i {
				set[j] = false
			}
		}
	}
	return count
}

// leetcode submit region end(Prohibit modification and deletion)
