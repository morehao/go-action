package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := createBinaryTree(0, arr)
	treeStr, _ := jsoniter.Marshal(tree)
	fmt.Println(string(treeStr))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(i int, nums []int) *TreeNode {
	if nums[i] == 0 {
		return nil
	}
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
