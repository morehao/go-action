package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	root := []int{3, 9, 20, 0, 0, 15, 7}
	tree := binaryTree.CreateBinaryTree(0, root)
	fmt.Println(levelOrder(tree))
}

// 层序遍历，BFS，返回结果为二维数组，奇数行正序访问，偶数行逆序访问
func levelOrder(root *binaryTree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		queue = []*binaryTree.TreeNode{root}
		res   [][]int
	)
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		var tempQueue []*binaryTree.TreeNode
		for j := 0; j < len(queue); j++ {
			currNode := queue[j]
			if (i+1)%2 == 1 {
				res[i] = append(res[i], currNode.Val)
			} else {
				res[i] = append(res[i], queue[len(queue)-1-j].Val)
			}
			if currNode.Left != nil {
				tempQueue = append(tempQueue, currNode.Left)
			}
			if currNode.Right != nil {
				tempQueue = append(tempQueue, currNode.Right)
			}
		}
		queue = tempQueue
	}
	return res
}
