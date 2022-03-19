package main

import "fmt"

func main() {
	nums := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := create(0, nums)
	fmt.Println(tree.LevelOrder())
	// fmt.Println(tree.InOrderTraversal())
	// fmt.Println(tree.PreOrderTraversal())
	// fmt.Println(tree.PostOrderTraversal())
	fmt.Println(TreeKth(tree, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var index int

func TreeKth(root *TreeNode, k int) int {
	if root == nil || k < 1 {
		return -1
	}
	index = 0
	node := convertSearch(root, k)
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

func create(i int, nums []int) *TreeNode {
	if nums[i] == 0 {
		return nil
	}
	node := &TreeNode{
		Val: nums[i],
	}
	if 2*i+1 < len(nums) {
		node.Left = create(2*i+1, nums)
	}
	if 2*i+2 < len(nums) {
		node.Right = create(2*i+2, nums)
	}
	return node
}

func (t *TreeNode) LevelOrder() [][]int {
	res := make([][]int, 0)
	q := []*TreeNode{t}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		nodeList := make([]*TreeNode, 0)
		for j := 0; j < len(q); j++ {
			if q[j].Left != nil {
				nodeList = append(nodeList, q[j].Left)
			}
			if q[j].Right != nil {
				nodeList = append(nodeList, q[j].Right)
			}
			res[i] = append(res[i], q[j].Val)
		}
		q = nodeList
	}
	return res
}

func (t *TreeNode) InOrderTraversal() []int {
	res := make([]int, 0)
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(t)
	return res
}

func (t *TreeNode) PreOrderTraversal() []int {
	res := make([]int, 0)
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(t)
	return res
}

func (t *TreeNode) PostOrderTraversal() []int {
	res := make([]int, 0)
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}
	postOrder(t)
	return res
}
