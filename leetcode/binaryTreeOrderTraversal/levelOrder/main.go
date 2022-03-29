package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	root := []int{3, 9, 20, 0, 0, 15, 7}
	tree := binaryTree.CreateBinaryTree(0, root)
	fmt.Println(levelOrder(tree))
}

func levelOrder(root *binaryTree.TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*binaryTree.TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*binaryTree.TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func levelOrderRightFirst(root *binaryTree.TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*binaryTree.TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*binaryTree.TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Right != nil {
				p = append(p, node.Right)
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
		}
		q = p
	}
	return ret
}
