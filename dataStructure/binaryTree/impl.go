package binaryTree

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

func (nd *TreeNode) Insert(newNode *TreeNode) {
	if newNode.Val == nd.Val {
		return
	}
	if newNode.Val > nd.Val {
		if nd.Right == nil {
			nd.Right = newNode
		} else {
			nd.Right.Insert(newNode)
		}
	} else {
		if nd.Left == nil {
			nd.Left = newNode
		} else {
			nd.Left.Insert(newNode)
		}
	}
}

func (nd *TreeNode) Search(value int) *TreeNode {
	if nd == nil {
		return nil
	}
	// 1、比较是否为当前节点
	if value == nd.Val {
		return nd
	}
	// 2、大于当前节点，递归右边
	if value > nd.Val {
		return nd.Right.Search(value)
	}
	// 3、大于当前节点，递归左边
	if value < nd.Val {
		return nd.Left.Search(value)
	}
	return nil
}

// Remove 删除一个元素.
func (nd *TreeNode) Remove(value int) bool {
	_, existed := remove(nd, value)
	return existed
}

// 用来递归移除节点的辅助方法.
// 返回替换root的新节点，以及元素是否存在
func remove(root *TreeNode, value int) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	// 从左边找
	if value < root.Val {
		root.Left, existed = remove(root.Left, value)
		return root, existed
	}
	// 从右边找
	if value > root.Val {
		root.Right, existed = remove(root.Right, value)
		return root, existed
	}
	// 如果此节点正是要移除的节点,那么返回此节点，同时返回之前可能需要调整.
	existed = true
	// 如果此节点没有孩子，直接返回即可
	if root.Left == nil && root.Right == nil {
		root = nil
		return root, existed
	}
	// 如果左子节点为空, 提升右子节点
	if root.Left == nil {
		root = root.Right
		return root, existed
	}
	// 如果右子节点为空, 提升左子节点
	if root.Right == nil {
		root = root.Left
		return root, existed
	}
	// 如果左右节点都存在,那么从右边节点找到一个最小的节点提升，这个节点肯定比左子树所有节点都大.
	// 也可以从左子树节点中找一个最大的提升，道理一样.
	smallestInRight := root.Right.Min()
	// 提升
	root.Val = smallestInRight
	// 从右边子树中移除此节点
	root.Right, _ = remove(root.Right, smallestInRight)
	return root, existed
}

func (nd *TreeNode) Min() int {
	if nd.Left == nil {
		return nd.Val
	}
	return nd.Left.Min()
}

func (nd *TreeNode) Max() int {
	if nd.Right == nil {
		return nd.Val
	}
	return nd.Right.Max()
}

// 层序遍历
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
