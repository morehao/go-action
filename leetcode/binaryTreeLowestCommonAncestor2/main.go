package binaryTreeLowestCommonAncestor2

import "go-practict/dataStructure/binaryTree"

func lowestCommonAncestor(root, p, q *binaryTree.TreeNode) *binaryTree.TreeNode {
	if root == nil {
		return nil
	}

	// x恰好是p节点或q 节点且它的左子树或右子树有一个包含了另一个节点的情况
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左子树和右子树均包含p 节点或q 节点，如果左子树包含的是p 节点，那么右子树只能包含q 节点，反之亦然，因为p节点和q节点都是不同且唯一的节点
	if left != nil && right != nil {
		return root
	}
	// left为空，说名p和q被右节点包含，右节点为祖先节点
	if left == nil {
		return right
	}
	return left
}
