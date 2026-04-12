/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 2080 👎 0

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

// BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})  // 初始化当前层的结果
		levelQueue := []*TreeNode{} // 初始化下一层队列
		for j := 0; j < len(queue); j++ {
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
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
