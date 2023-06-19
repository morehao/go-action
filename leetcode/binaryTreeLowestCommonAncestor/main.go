package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := binaryTree.ArrayToTree(nums)
	fmt.Println(root.LevelOrder())
	fmt.Println(lowestCommonAncestor(nil, nil, nil))
}

func lowestCommonAncestor(root, p, q *binaryTree.TreeNode) *binaryTree.TreeNode {
	ancestor := root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return ancestor
		}
	}
}
