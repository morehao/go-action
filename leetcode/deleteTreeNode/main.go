package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 7}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(deleteNode(tree, 3))
	fmt.Println(tree.LevelOrder())
}

func deleteNode(root *binaryTree.TreeNode, key int) *binaryTree.TreeNode {
	// 递归删除节点
	// 5-case
	// 1.空节点返回nil
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 2.左右孩子都空，直接删除节点，返回nil
		if root.Left == nil && root.Right == nil {
			root = nil
			return nil
		}
		// 3.左空右不空，删除节点，右孩子补位
		if root.Left == nil && root.Right != nil {
			root = root.Right
			return root
		}
		// 4.左不空右空，删除节点，左孩子补位
		if root.Left != nil && root.Right == nil {
			root = root.Left
			return root
		}
		// 5.左右都不空时，右孩子补位，将删除节点的左子树，
		// 放到删除节点的右子树的最左面节点的左孩子位置
		left := root.Left
		right := root.Right
		tmp := root.Right
		// 通过tmp寻找root.Right的最左孩子节点
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 右子树的最小节点，该节点大于root.Left的整颗子树，Left指向left子树
		tmp.Left = left
		root = right // 右孩子补位
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
