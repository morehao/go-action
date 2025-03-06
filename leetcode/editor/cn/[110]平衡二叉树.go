/**
给定一个二叉树，判断它是否是 平衡二叉树



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：true


 示例 2：


输入：root = [1,2,2,3,3,null,null,4,4]
输出：false


 示例 3：


输入：root = []
输出：true




 提示：


 树中的节点数在范围 [0, 5000] 内
 -10⁴ <= Node.val <= 10⁴


 Related Topics 树 深度优先搜索 二叉树 👍 1582 👎 0

*/

package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	depthDiff := abs(maxDepth(root.Left) - maxDepth(root.Right))
	return depthDiff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// leetcode submit region end(Prohibit modification and deletion)
