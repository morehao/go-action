package main

import (
	"fmt"

	"go-practict/interview/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(preorderTraversal(tree))
}

// 先序遍历：优先于后代节点的顺序访问每个节点,即先遍历当前节点然后左节点再右节点
// 应用场景：打印结构化的文档
func preorderTraversal(root *binaryTree.TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *binaryTree.TreeNode)
	preorder = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}
