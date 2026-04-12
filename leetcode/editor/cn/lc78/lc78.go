/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



 示例 1：


输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]


 示例 2：


输入：nums = [0]
输出：[[],[0]]




 提示：


 1 <= nums.length <= 10
 -10 <= nums[i] <= 10
 nums 中的所有元素 互不相同


 Related Topics 位运算 数组 回溯 👍 2454 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	// 初始子集为空集
	result := make([][]int, 1)

	for _, num := range nums {
		n := len(result)
		// 遍历 res 中已有的所有子集，将当前元素放入到子集中
		for i := 0; i < n; i++ {
			// 复制当前子集并加入新元素
			newSubset := make([]int, len(result[i]))
			copy(newSubset, result[i])
			newSubset = append(newSubset, num)
			result = append(result, newSubset)
		}
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)
