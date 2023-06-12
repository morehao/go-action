package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{1, 0, 2, 3}
	root := binaryTree.CreateBinaryTree(0, nums)
	fmt.Println(root)
	fmt.Println(inorderTraversal(root))
}

// 中序遍历，先遍历其左子树，然后访问根，最后遍历其右子树
func inorderTraversal(root *binaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}
	var fn func(node *binaryTree.TreeNode)
	var res []int
	fn = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		res = append(res, node.Val)
		fn(node.Right)
	}
	fn(root)
	return res
}
