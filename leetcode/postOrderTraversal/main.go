package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(postorderTraversal(tree))
}

// 后序遍历：先访问后代节点再访问本身节点
// 应用场景：计算一个目录和它的子目录中所有文件所占空间的大小
func postorderTraversal(root *binaryTree.TreeNode) []int {
	res := make([]int, 0)
	var postorder func(node *binaryTree.TreeNode)
	postorder = func(node *binaryTree.TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}
