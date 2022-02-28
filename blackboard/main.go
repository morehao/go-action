package main

import (
	"fmt"

	"go-practict/interview/binaryTree"
)

func main() {
	arr := []int{3, 9, 20, 0, 0, 15, 7}
	root := binaryTree.CreateBinaryTree(0, arr)

	fmt.Println(levelOrderPrint(root))
}

func levelOrderPrint(root *binaryTree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodeList := []*binaryTree.TreeNode{root}
	for i := 0; len(nodeList) > 0; i++ {
		res = append(res, []int{})
		var currentLevelNodes []*binaryTree.TreeNode
		for j := 0; j < len(nodeList); j++ {
			node := nodeList[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				currentLevelNodes = append(currentLevelNodes, node.Left)
			}
			if node.Right != nil {
				currentLevelNodes = append(currentLevelNodes, node.Right)
			}
		}
		nodeList = currentLevelNodes
	}
	return res
}
