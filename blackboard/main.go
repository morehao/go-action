package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := createTree(0, arr)
	fmt.Println(tree.LevelOrder())
	fmt.Println(InOrder(tree))
}

func createTree(i int, nums []int) *binaryTree.TreeNode {
	if nums[i] == 0 {
		return nil
	}
	root := &binaryTree.TreeNode{
		Val: nums[i],
	}
	if i < len(nums) && 2*i+1 < len(nums) {
		root.Left = createTree(2*i+1, nums)
	}
	if i < len(nums) && 2*i+2 < len(nums) {
		root.Right = createTree(2*i+2, nums)
	}
	return root
}

func levelOrder(root *binaryTree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*binaryTree.TreeNode{root}
	res := make([][]int, 0)
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		current := make([]*binaryTree.TreeNode, 0)
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				current = append(current, node.Left)
			}
			if node.Right != nil {
				current = append(current, node.Right)
			}
		}
		q = current
	}
	return res
}

func InOrder(root *binaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	var inorder func(node *binaryTree.TreeNode)
	inorder = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		inorder(node.Left)
		inorder(node.Right)
	}
	inorder(root)
	return res
}
