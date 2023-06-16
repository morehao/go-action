package main

import "go-practict/dataStructure/binaryTree"

func isSymmetric(root *binaryTree.TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(a, b *binaryTree.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return check(a.Left, b.Right) && check(a.Right, b.Left)
}
