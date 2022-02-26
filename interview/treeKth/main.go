package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 0, 2}
	tree := createBinaryTree(0, arr)
	fmt.Println(kthLargest(tree, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var index int

func kthLargest(root *TreeNode, k int) int {
	if root == nil || k < 1 {
		return -1
	}
	index = 0
	node := convertToSearch(root, k)
	return node.Val
}

func convertToSearch(node *TreeNode, k int) *TreeNode {
	if node.Right != nil {
		right := convertToSearch(node.Right, k)
		if right != nil {
			return right
		}
	}

	index++
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

func createBinaryTree(i int, nums []int) *TreeNode {
	tree := &TreeNode{nums[i], nil, nil}
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
