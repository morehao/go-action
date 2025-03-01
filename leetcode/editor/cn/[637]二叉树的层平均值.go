/**
给定一个非空二叉树的根节点
 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10⁻⁵ 以内的答案可以被接受。



 示例 1：




输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。


 示例 2:




输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]




 提示：





 树中节点数量在 [1, 10⁴] 范围内
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 520 👎 0

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

type LevelItem struct {
	Sum   int
	Count int
}

func averageOfLevels(root *TreeNode) []float64 {
	levelList := make([]LevelItem, 0)
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 当前 level 已经在 levelList 中
		if level < len(levelList) {
			levelList[level].Sum += node.Val
			levelList[level].Count++
		} else {
			levelList = append(levelList, LevelItem{
				Sum:   node.Val,
				Count: 1,
			})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	avgList := make([]float64, len(levelList))
	for i := range levelList {
		avgList[i] = float64(levelList[i].Sum) / float64(levelList[i].Count)
	}
	return avgList
}

// leetcode submit region end(Prohibit modification and deletion)
