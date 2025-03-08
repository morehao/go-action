/**
待传输文件被切分成多个部分，按照原排列顺序，每部分文件编号均为一个 正整数（至少含有两个文件）。传输要求为：连续文件编号总和为接收方指定数字 target 的所
有文件。请返回所有符合该要求的文件传输组合列表。

 注意，返回时需遵循以下规则：


 每种组合按照文件编号 升序 排列；
 不同组合按照第一个文件编号 升序 排列。




 示例 1：


输入：target = 12
输出：[[3, 4, 5]]
解释：在上述示例中，存在一个连续正整数序列的和为 12，为 [3, 4, 5]。


 示例 2：


输入：target = 18
输出：[[3,4,5,6],[5,6,7]]
解释：在上述示例中，存在两个连续正整数序列的和分别为 18，分别为 [3, 4, 5, 6] 和 [5, 6, 7]。




 提示：


 1 <= target <= 10^5




 Related Topics 数学 双指针 枚举 👍 583 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func fileCombination(target int) [][]int {
	var result [][]int
	left, right := 1, 1 // 滑动窗口的左右边界
	sum := 0            // 当前窗口内文件编号的总和

	for left <= target/2 {
		if sum < target {
			// 如果总和小于 target，扩大窗口（右边界右移）
			sum += right
			right++
		} else if sum > target {
			// 如果总和大于 target，缩小窗口（左边界右移）
			sum -= left
			left++
		} else {
			// 如果总和等于 target，记录当前窗口内的文件编号
			var sequence []int
			for i := left; i < right; i++ {
				sequence = append(sequence, i)
			}
			result = append(result, sequence)
			// 继续寻找下一个组合
			sum -= left
			left++
		}
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)
