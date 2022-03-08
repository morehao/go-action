package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(inorderTraversal(tree))
}

// 二叉搜索树节点值按从小到大遍历
// 应用场景：二叉搜索树排序
func inorderTraversal(root *binaryTree.TreeNode) []int {
	res := make([]int, 0)
	var inorder func(node *binaryTree.TreeNode)
	inorder = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}
