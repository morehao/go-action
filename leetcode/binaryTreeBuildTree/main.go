package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
)

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root.LevelOrder())
}

/*
递归思路
对于任意一颗树而言，前序遍历的形式总是:[根节点,[左子树的前序遍历结果],[右子树的前序遍历结果]],
即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是:[[左子树的中序遍历结果],根节点,[右子树的中序遍历结果]]
*/
func buildTree(preorder []int, inorder []int) *binaryTree.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &binaryTree.TreeNode{Val: preorder[0]}
	inorderRootIndex := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			inorderRootIndex = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:inorderRootIndex])+1], inorder[:inorderRootIndex])
	root.Right = buildTree(preorder[len(inorder[:inorderRootIndex])+1:], inorder[inorderRootIndex+1:])
	return root
}
