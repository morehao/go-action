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

func levelOrder(root *binaryTree.TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	q := []*binaryTree.TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		ret = append(ret, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return ret
}
