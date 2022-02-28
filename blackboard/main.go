package main

import (
	"fmt"

	"go-practict/interview/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 7}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(deleteNode(tree, 3))
	fmt.Println(tree.LevelOrder())
}

func deleteNode(root *binaryTree.TreeNode, val int) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = deleteNode(root.Left, val)
		return root
	}
	if val > root.Val {
		root.Right = deleteNode(root.Right, val)
		return root
	}
	if root.Left == nil && root.Right == nil {
		root = nil
		return root
	}
	if root.Left == nil && root.Right != nil {
		root = root.Left
		return root
	}
	if root.Left != nil && root.Right == nil {
		return root
	}
	left := root.Left
	right := root.Right
	tmp := root.Right
	for tmp.Left != nil {
		tmp = tmp.Left
	}
	tmp.Left = left
	root = right
	return root
}
