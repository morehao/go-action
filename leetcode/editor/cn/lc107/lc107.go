/**
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[15,7],[9,20],[3]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 831 👎 0

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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		n := len(queue)
		var levelQueue []*TreeNode
		for j := 0; j < n; j++ {
			node := queue[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				levelQueue = append(levelQueue, node.Left)
			}
			if node.Right != nil {
				levelQueue = append(levelQueue, node.Right)
			}
		}
		queue = levelQueue
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
