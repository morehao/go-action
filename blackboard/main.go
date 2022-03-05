package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 7}
	tree := createBinaryTree(0, arr)
	deleteNode(tree, 3)
	fmt.Println(tree.levelOrder())

}

func deleteNode(root *treeNode, val int) *treeNode {
	if root == nil {
		return nil
	}
	if root.val > val {
		root.left = deleteNode(root.left, val)
	}
	if root.val < val {
		root.right = deleteNode(root.right, val)
	}
	if root.val == val {
		if root.left == nil && root.right == nil {
			root = nil
			return root
		}
		if root.left == nil && root.right != nil {
			root = root.right
			return root
		}
		if root.left != nil && root.right == nil {
			root = root.left
			return root
		}
		if root.left != nil && root.right != nil {
			left, right := root.left, root.right
			tmp := root.right
			for tmp.left != nil {
				tmp = tmp.left
			}
			tmp.left = left
			root = right
			return root
		}
	}
	return root
}

func createBinaryTree(i int, nums []int) *treeNode {
	if nums[i] == 0 {
		return nil
	}
	tree := &treeNode{
		val: nums[i],
	}
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.left = createBinaryTree(2*i+1, nums)
	}
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.right = createBinaryTree(2*i+2, nums)
	}
	return tree
}

func (t *treeNode) levelOrder() [][]int {
	nodeList := make([]*treeNode, 0)
	nodeList = append(nodeList, t)
	res := make([][]int, 0)
	for i := 0; len(nodeList) > 0; i++ {
		res = append(res, []int{})
		var currentNodeList []*treeNode
		for j := 0; j < len(nodeList); j++ {
			node := nodeList[j]
			res[i] = append(res[i], node.val)
			if node.left != nil {
				currentNodeList = append(currentNodeList, node.left)
			}
			if node.right != nil {
				currentNodeList = append(currentNodeList, node.right)
			}
		}
		nodeList = currentNodeList
	}
	return res
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}
