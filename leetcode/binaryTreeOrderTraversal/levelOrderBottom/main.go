package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	nums := []int{3, 9, 20, 0, 0, 15, 7}
	tree := binaryTree.CreateBinaryTree(0, nums)
	fmt.Println(levelOrderBottom(tree))
}

func levelOrderBottom(root *binaryTree.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*binaryTree.TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
