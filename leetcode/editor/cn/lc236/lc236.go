/**
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个
节点也可以是它自己的祖先）。”



 示例 1：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。


 示例 2：


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。


 示例 3：


输入：root = [1,2], p = 1, q = 2
输出：1




 提示：


 树中节点数目在范围 [2, 10⁵] 内。
 -10⁹ <= Node.val <= 10⁹
 所有 Node.val 互不相同 。
 p != q
 p 和 q 均存在于给定的二叉树中。


 Related Topics 树 深度优先搜索 二叉树 👍 2906 👎 0

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果当前节点就是 p 或 q，它可能是公共祖先之一，因此直接返回当前节点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果 left 和 right 都不为空，说明 p 和 q 分别位于当前节点的左右子树中，即它们的最近公共祖先是当前节点 root。
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// leetcode submit region end(Prohibit modification and deletion)
