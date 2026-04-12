/**
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。 

 

 示例 1： 

 
 输入：root = [1,null,2,3] 
 

 输出：[1,2,3] 

 解释： 

 

 示例 2： 

 
 输入：root = [1,2,3,4,5,null,8,null,null,6,7,9] 
 

 输出：[1,2,4,5,6,7,3,8,9] 

 解释： 

 

 示例 3： 

 
 输入：root = [] 
 

 输出：[] 

 示例 4： 

 
 输入：root = [1] 
 

 输出：[1] 

 

 提示： 

 
 树中节点数目在范围 [0, 100] 内 
 -100 <= Node.val <= 100 
 

 

 进阶：递归算法很简单，你可以通过迭代算法完成吗？ 

 Related Topics 栈 树 深度优先搜索 二叉树 👍 1317 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		fn(node.Left)
		fn(node.Right)
	}
	fn(root)
	return res
    
}
//leetcode submit region end(Prohibit modification and deletion)
