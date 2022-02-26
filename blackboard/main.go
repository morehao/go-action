package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	root := createBinaryTree(0, arr)

	fmt.Println(treeLast(root, 3))
}

func treeLast(root *TreeNode, k int) int {
	if k < 1 || root == nil {
		return -1
	}
	index = 0
	node := search(root, k)
	return node.Value
}

var index int

func search(node *TreeNode, k int) *TreeNode {
	if node.Right != nil {
		right := search(node.Right, k)
		if right != nil {
			return right
		}
	}
	index++
	if index == k {
		return node
	}
	if node.Left != nil {
		right := search(node.Left, k)
		if right != nil {
			return right
		}
	}
	return nil
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(i int, nums []int) *TreeNode {
	if nums[i] == 0 {
		return nil
	}
	treeNode := &TreeNode{
		Value: nums[i],
	}
	if i < len(nums) && 2*i+1 < len(nums) {
		treeNode.Left = createBinaryTree(2*i+1, nums)
	}
	if i < len(nums) && 2*i+2 < len(nums) {
		treeNode.Right = createBinaryTree(2*i+2, nums)
	}
	return treeNode
}
