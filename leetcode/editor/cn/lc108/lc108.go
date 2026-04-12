/**
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。



 示例 1：


输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：



 示例 2：


输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。




 提示：


 1 <= nums.length <= 10⁴
 -10⁴ <= nums[i] <= 10⁴
 nums 按 严格递增 顺序排列


 Related Topics 树 二叉搜索树 数组 分治 二叉树 👍 1619 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return fn(nums, 0, len(nums)-1)
}
func fn(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = fn(nums, left, mid-1)
	root.Right = fn(nums, mid+1, right)
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
