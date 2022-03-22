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
		right := convertSearch(node.Right, k)
		if right != nil {
			return right
		}
	}
	index++
	if index == k {
		return node
	}
	if node.Left != nil {
		left := convertSearch(node.Left, k)
		if left != nil {
			return left
		}
	}
	return nil
}
