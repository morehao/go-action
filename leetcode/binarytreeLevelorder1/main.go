package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []int{1, 0, 2, 3}
	root := binaryTree.CreateBinaryTree(0, nums)
	fmt.Println(root)
	fmt.Println(levelOrder(root))
}

// 层序遍历，BFS，输出结果为一维数组
func levelOrder(root *binaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res   []int
		queue []*binaryTree.TreeNode
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		res = append(res, currNode.Val)
		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}
	}
	return res
}

func levelOrder2(root *binaryTree.TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res   []int
		queue []*binaryTree.TreeNode
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		var tempQueue []*binaryTree.TreeNode
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			res = append(res, node.Val)
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}
		queue = tempQueue
	}
	return res
}
