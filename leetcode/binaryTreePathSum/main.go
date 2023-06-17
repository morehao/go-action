package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}
	root := binaryTree.BuildTreeWithNums(nums)
	fmt.Println(pathSum(root, 22))
}

// 深度优先算法DFS
func pathSum(root *binaryTree.TreeNode, target int) [][]int {
	var (
		res [][]int
		fn  func(node *binaryTree.TreeNode, sum int, path []int)
	)
	fn = func(node *binaryTree.TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		if sum == 0 && node.Left == nil && node.Right == nil {
			// path在改变，使用新的底层数组
			res = append(res, append([]int{}, path...))
			return
		}
		fn(node.Left, sum, path)
		fn(node.Right, sum, path)
	}
	fn(root, target, []int{})
	return res
}

// 回溯+DFS
func pathSum2(root *binaryTree.TreeNode, target int) [][]int {
	var (
		path []int
		res  [][]int
		dfs  func(node *binaryTree.TreeNode, sum int)
	)
	dfs = func(node *binaryTree.TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		// 路径恢复：向上回溯前，要删除当前节点
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return res
}
