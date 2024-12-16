package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{1, 0, 2, 3}
	root := binaryTree.CreateBinaryTree(0, nums)
	fmt.Println(root)
	fmt.Println(postorderTraversal(root))
}

// 后序遍历，先遍历其左子树，然后遍历其右子树，最后访问根
func postorderTraversal(root *binaryTree.TreeNode) []int {
	var fn func(node *binaryTree.TreeNode)
	var res []int
	fn = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		fn(node.Right)
		res = append(res, node.Val)
	}
	fn(root)
	return res
}
