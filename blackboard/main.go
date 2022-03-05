package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 7}
	tree := createBinaryTree(0, arr)
	fmt.Println(tree.levelOrder())

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
