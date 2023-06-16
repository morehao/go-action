package main

import "go-practict/dataStructure/binaryTree"

func mirrorTree(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
