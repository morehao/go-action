package main

import "go-practict/dataStructure/binaryTree"

func mirrorTree(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
