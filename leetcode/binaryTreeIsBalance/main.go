package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := binaryTree.ArrayToTree(nums)
	fmt.Println(root.LevelOrder())
	fmt.Println(isBalanced(root))
}

func isBalanced(root *binaryTree.TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *binaryTree.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
