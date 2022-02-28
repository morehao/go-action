package main

import (
	"fmt"

	"go-practict/interview/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(tree.LevelOrder())
	fmt.Println(kthLargest(tree, 3))
}

func kthLargest(root *binaryTree.TreeNode, k int) int {
	if root == nil || k < 1 {
		return -1
	}
	index = 0
	node := search(root, k)
	return node.Val
}

var index int

func search(root *binaryTree.TreeNode, k int) *binaryTree.TreeNode {
	if root.Right != nil {
		right := search(root.Right, k)
		if right != nil {
			return right
		}
	}
	index++
	if k == index {
		return root
	}
	if root.Left != nil {
		left := search(root.Left, k)
		if left != nil {
			return left
		}
	}
	return nil
}
