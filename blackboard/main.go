package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建二叉树
func CreateBinaryTree(i int, nums []int) *TreeNode {
	if len(nums) == 0 || nums[i] == 0 {
		return nil
	}
	tree := &TreeNode{
		Val: nums[i],
	}
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = CreateBinaryTree(2*i+1, nums)
	}
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Left = CreateBinaryTree(2*i+2, nums)
	}
	return tree
}

func PreLevelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		fn(node.Left)
		fn(node.Right)
	}
	fn(root)
	return res
}
