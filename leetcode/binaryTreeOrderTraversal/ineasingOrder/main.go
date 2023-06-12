package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{5, 3, 6, 2, 4, 0, 8, 1, 0, 0, 0, 7, 9}
	tree := binaryTree.CreateBinaryTree(0, nums)
	newTree := increasingOrder(tree)
	fmt.Println(newTree.LevelOrder())
}

func increasingOrder(root *binaryTree.TreeNode) *binaryTree.TreeNode {
	nums := make([]int, 0)
	var inOrder func(node *binaryTree.TreeNode)
	inOrder = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	dummyNode := &binaryTree.TreeNode{}
	currentNode := dummyNode
	for _, v := range nums {
		currentNode.Right = &binaryTree.TreeNode{
			Val: v,
		}
		currentNode = currentNode.Right
	}
	return dummyNode.Right
}
