package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}
	root := binaryTree.BuildTreeWithNums(nums)
	fmt.Println(root)
	fmt.Println(inorderTraversal(root))
}

// 中序遍历，先遍历其左子树，然后访问根，最后遍历其右子树
func inorderTraversal(root *binaryTree.TreeNode) []int {
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
