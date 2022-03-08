package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(treeKth(tree, 3))
}

func treeKth(root *binaryTree.TreeNode, k int) int {
	if root == nil || k < 1 {
		return -1
	}
	index = 0
	node := convertToSearch(root, k)
	return node.Val
}

var index int

func convertToSearch(node *binaryTree.TreeNode, k int) *binaryTree.TreeNode {
	if node.Right != nil {
		right := convertToSearch(node.Right, k)
		if right != nil {
			return right
		}

	}
	index++
	if index == k {
		return node
	}
	if node.Left != nil {
		left := convertToSearch(node.Left, k)
		if left != nil {
			return left
		}
	}
	return nil
}
