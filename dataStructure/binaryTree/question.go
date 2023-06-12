package binaryTree

var index int

func (nd *TreeNode) KthLargest(k int) int {
	if nd == nil || k < 1 {
		return -1
	}
	index = 0
	node := convertSearch(nd, k)
	return node.Val
}
func convertSearch(node *TreeNode, k int) *TreeNode {
	if node.Right != nil {
		Right := convertSearch(node.Right, k)
		if Right != nil {
			return Right
		}
	}
	index++
	if index == k {
		return node
	}
	if node.Left != nil {
		Left := convertSearch(node.Left, k)
		if Left != nil {
			return Left
		}
	}
	return nil
}
