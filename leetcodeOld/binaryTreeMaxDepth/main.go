package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := binaryTree.BuildTreeWithArray(nums)
	fmt.Println(root.LevelOrder())
	fmt.Println(maxDepth(root))
	fmt.Println(maxDepthBfs(root))
}

// dfs递归，最大深度为max(left, right) + 1
func maxDepth(root *binaryTree.TreeNode) int {
	if root == nil {
		return 0
	}
	defaultDepth := 1
	return max(maxDepth(root.Left), maxDepth(root.Right)) + defaultDepth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// bfs,每遍历一层，则计数器+1 ，直到遍历完成，则可得到树的深度。
func maxDepthBfs(root *binaryTree.TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		depth = 0
		queue = []*binaryTree.TreeNode{root}
	)
	for len(queue) > 0 {
		queueLen := len(queue)
		for queueLen > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queueLen--
		}
		depth++
	}
	return depth
}
