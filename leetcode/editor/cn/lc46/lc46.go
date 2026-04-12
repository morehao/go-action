/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



 示例 1：


输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


 示例 2：


输入：nums = [0,1]
输出：[[0,1],[1,0]]


 示例 3：


输入：nums = [1]
输出：[[1]]




 提示：


 1 <= nums.length <= 6
 -10 <= nums[i] <= 10
 nums 中的所有整数 互不相同


 Related Topics 数组 回溯 👍 3068 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	// 定义结果集 res，用于存储所有的排列结果
	var res [][]int

	// 获取输入数组 nums 的长度
	n := len(nums)

	// 创建一个大小为 n 的 curr 数组，用于存储当前排列的状态
	// 创建一个大小为 n 的 visited 数组，用于标记每个数字是否已经在当前排列中
	curr := make([]int, n)
	visited := make([]bool, n)

	// 定义递归函数 dfs，i 表示当前排列的索引位置
	var dfs func(int)
	dfs = func(i int) {
		// 如果当前索引等于数组长度，说明已经找到了一个完整的排列
		if i == n {
			// 将当前排列的 curr 加入结果集 res，使用 append([]int(nil), curr...) 来复制 curr 的内容
			res = append(res, append([]int{}, curr...))
			return
		}

		// 遍历每个数字，尝试将它放入当前位置 i
		// visited[j] 用于检查该数字是否已经在当前排列中
		for j, on := range visited {
			// 如果 nums[j] 还没有被放入排列中（即 visited[j] 为 false）
			if !on {
				// 将 nums[j] 放入当前位置 i
				curr[i] = nums[j]
				// 标记 nums[j] 已经被使用
				visited[j] = true
				// 递归调用，处理下一个位置
				dfs(i + 1)
				// 回溯，将 nums[j] 标记为未使用，恢复状态
				visited[j] = false
			}
		}
	}

	// 从索引 0 开始进行递归
	dfs(0)

	// 返回所有排列结果
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
