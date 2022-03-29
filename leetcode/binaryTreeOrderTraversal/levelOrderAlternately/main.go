package main

import (
	"fmt"

	"go-practict/leetcode/binaryTree"
)

func main() {
	root := []int{3, 9, 20, 0, 0, 15, 7}
	tree := binaryTree.CreateBinaryTree(0, root)
	fmt.Println(levelOrder(tree))
}

func levelOrder(root *binaryTree.TreeNode) [][]int {
	var res = make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*binaryTree.TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, length)
		for i := 0; i < length; i++ {
			if queue[i] == nil {
				continue
			}
			if len(res)%2 == 0 {
				// 偶数放队头
				tmp[i] = queue[i].Val
			} else {
				// 奇数放队尾
				tmp[length-1-i] = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
		res = append(res, tmp)
	}
	return res
}
