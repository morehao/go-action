package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(tree.LevelOrder())
	fmt.Println(kthLargest(tree, 3))
}

var index int

func kthLargest(root *binaryTree.TreeNode, k int) int {
	if root == nil || k < 1 {
		return -1
	}
	index = 0
	// 第k大其实就是二叉树第k层的右节点
	node := convertToSearch(root, k)
	return node.Val
}

func convertToSearch(node *binaryTree.TreeNode, k int) *binaryTree.TreeNode {
	if node.Right != nil {
		right := convertToSearch(node.Right, k)
		if right != nil {
			return right
		}
	}

	index++
	// 第k大在第k层最右节点
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
