package main

import (
	"go-practict/dataStructure/binaryTree"
)

// 中序遍历结果为递增序列
func kthLargest(root *binaryTree.TreeNode, k int) int {
	var (
		nums    []int
		orderFn func(node *binaryTree.TreeNode)
	)
	orderFn = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		orderFn(node.Left)
		nums = append(nums, node.Val)
		orderFn(node.Right)
	}
	orderFn(root)
	// 取倒数的元素
	return nums[len(nums)-k]
}

// 中序遍历的倒序结果为递增序列
func kthLargest2(root *binaryTree.TreeNode, k int) int {
	ans := 0
	var dfs func(node *binaryTree.TreeNode)
	dfs = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		k--
		if k == 0 {
			ans = node.Val
		}
		dfs(node.Left)
	}
	dfs(root)
	return ans
}
