/**
某公司组织架构以二叉搜索树形式记录，节点值为处于该职位的员工编号。请返回第 cnt 大的员工编号。



 示例 1：




输入：root = [7, 3, 9, 1, 5], cnt = 2
       7
      / \
     3   9
    / \
   1   5
输出：7


 示例 2：




输入: root = [10, 5, 15, 2, 7, null, 20, 1, null, 6, 8], cnt = 4
       10
      / \
     5   15
    / \    \
   2   7    20
  /   / \
 1   6   8
输出: 8



 提示：


 1 ≤ cnt ≤ 二叉搜索树元素个数




 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 423 👎 0

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
func findTargetNode(root *TreeNode, cnt int) int {
	var result int
	var count int

	// 逆中序遍历（右-根-左）
	var reverseInorder func(node *TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 遍历右子树
		reverseInorder(node.Right)

		// 处理当前节点
		count++
		if count == cnt {
			result = node.Val
			return
		}

		// 遍历左子树
		reverseInorder(node.Left)
	}

	// 开始遍历
	reverseInorder(root)
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
