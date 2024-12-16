package main

import (
	"go-practict/dataStructure/binaryTree"
)

func isSubStructure(A *binaryTree.TreeNode, B *binaryTree.TreeNode) bool {
	// 1、都为nil，直接返回false
	if A == nil || B == nil {
		return false
	}
	// 2、判断是否为完全相同的二叉树
	if isSameNode(A, B) {
		return true
	}
	// 3、是否为包含关系
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSameNode(a, b *binaryTree.TreeNode) bool {
	// b为nil，说明二叉树b已经遍历结束，b是a的子结构
	if b == nil {
		return true
	}
	// a为nil，说明二叉树a已经遍历结束，b不是a的子结构
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	// 递归遍历左右节点再进行比较
	return isSameNode(a.Left, b.Left) && isSameNode(a.Right, b.Right)
}
