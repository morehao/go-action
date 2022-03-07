package main

import (
	"fmt"

	"go-practict/interview/binaryTree"
)

func main() {
	root := []int{3, 9, 20, 0, 0, 15, 7}
	tree := binaryTree.CreateBinaryTree(0, root)
	fmt.Println(levelOrder(tree))
}

func levelOrder(root *binaryTree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodeList := []*binaryTree.TreeNode{root}
	res := make([][]int, 0)
	for i := 0; len(nodeList) > 0; i++ {
		res = append(res, []int{})
		currentNodeList := make([]*binaryTree.TreeNode, 0)
		for j := 0; j < len(nodeList); j++ {
			node := nodeList[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				currentNodeList = append(currentNodeList, node.Left)
			}
			if node.Right != nil {
				currentNodeList = append(currentNodeList, node.Right)
			}
		}
		nodeList = currentNodeList
	}
	return res
}
