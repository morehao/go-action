package main

import (
	"fmt"

	"go-practict/interview/binaryTree"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := binaryTree.CreateBinaryTree(0, arr)
	fmt.Println(kthLargest(tree, 3))
}

type TreeNode struct {
	Val   int
	Left  *binaryTree.TreeNode
	Right *binaryTree.TreeNode
}

var index int

func kthLargest(root *binaryTree.TreeNode, k int) int {
	if root == nil || k < 1 {
		return -1
	}
	index = 0
	node := convertToSearch(root, k)
	return node.Val
}

func convertToSearch(node *binaryTree.TreeNode, k int) *binaryTree.TreeNode {
	fmt.Println(node)
	if node.Right != nil {
		right := convertToSearch(node.Right, k)
		if right != nil {
			return right
		}
	}

	index++
	// 第k大在第k层最右节点
	if index == k {
		return node
	}

	if node.Left != nil {
		left := convertToSearch(node.Left, k)
		if left != nil {
			return left
		}
	}

	return nil
}

func createBinaryTree(i int, nums []int) *binaryTree.TreeNode {
	if nums[i] == 0 {
		return nil
	}
	tree := &binaryTree.TreeNode{nums[i], nil, nil}
	// 左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = createBinaryTree(2*i+1, nums)
	}
	// 右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Right = createBinaryTree(2*i+2, nums)
	}
	return tree
}
