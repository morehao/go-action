package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{1, 0, 2, 3}
	root := binaryTree.CreateBinaryTree(0, nums)
	fmt.Println(root)
	fmt.Println(preorderTraversal(root))
}

// 二叉树前序遍历，先访问根，然后遍历其左子树，最后遍历其右子树
func preorderTraversal(root *binaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var fn func(node *binaryTree.TreeNode)
	fn = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		fn(node.Left)
		fn(node.Right)
	}
	fn(root)
	return res
}
