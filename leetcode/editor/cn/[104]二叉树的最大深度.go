/**
给定一个二叉树 root ，返回其最大深度。

 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。



 示例 1：






输入：root = [3,9,20,null,null,15,7]
输出：3


 示例 2：


输入：root = [1,null,2]
输出：2




 提示：


 树中节点的数量在 [0, 10⁴] 区间内。
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1928 👎 0

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
// 深度优先搜索
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// leetcode submit region end(Prohibit modification and deletion)
