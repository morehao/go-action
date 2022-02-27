package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBinaryTree(i int, nums []int) *TreeNode {
	if nums[i] == 0 {
		return nil
	}
	tree := &TreeNode{nums[i], nil, nil}
	// 左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = CreateBinaryTree(2*i+1, nums)
	}
	// 右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Right = CreateBinaryTree(2*i+2, nums)
	}
	return tree
}

func (nd *TreeNode) LevelOrder() [][]int {
	var ret [][]int
	if nd == nil {
		return ret
	}
	q := []*TreeNode{nd}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
