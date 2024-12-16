package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := binaryTree.BuildTreeWithArray(nums)
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
			// 如果当前节点的值不满足上述两条要求，那么说明当前节点就是「分岔点」。此时，p和q要么在当前节点的不同的子树中，要么其中一个就是当前节点。
			return ancestor
		}
	}
}
