/**
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所
有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

 candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

 对于给定的输入，保证和为 target 的不同组合数少于 150 个。



 示例 1：


输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

 示例 2：


输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

 示例 3：


输入: candidates = [2], target = 1
输出: []




 提示：


 1 <= candidates.length <= 30
 2 <= candidates[i] <= 40
 candidates 的所有元素 互不相同
 1 <= target <= 40


 Related Topics 数组 回溯 👍 2997 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	// 结果集，用来存储所有符合条件的组合
	var res [][]int

	// 当前组合的数组
	var comb []int

	// 递归函数，负责从当前位置开始尝试组合
	var fn func(target, idx int)
	fn = func(target, idx int) {
		// 如果已遍历完所有候选数字，则返回
		if idx == len(candidates) {
			return
		}

		// 如果目标值为 0，说明当前组合成功，添加到结果集
		if target == 0 {
			// 通过 append([]int{}, comb...) 创建一个当前组合的副本，避免后续修改
			res = append(res, append([]int{}, comb...))
			return
		}

		// 不选择当前候选数字，递归进入下一个候选数字
		fn(target, idx+1)

		// 如果当前数字小于等于剩余的目标值，则可以继续选择当前数字
		if target-candidates[idx] >= 0 {
			// 选择当前数字，加入当前组合
			comb = append(comb, candidates[idx])

			// 继续递归，保持当前位置 (idx) 不变，可以重复选取当前数字
			fn(target-candidates[idx], idx)

			// 回溯，移除当前选择的数字，恢复当前组合状态
			comb = comb[:len(comb)-1]
		}
	}

	// 从目标值和第一个候选数字开始递归
	fn(target, 0)

	// 返回结果集
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
