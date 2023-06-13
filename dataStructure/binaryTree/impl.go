package binaryTree

func (t *TreeNode) Insert(newNode *TreeNode) {
	if newNode.Val == t.Val {
		return
	}
	if newNode.Val > t.Val {
		if t.Right == nil {
			t.Right = newNode
		} else {
			t.Right.Insert(newNode)
		}
	} else {
		if t.Left == nil {
			t.Left = newNode
		} else {
			t.Left.Insert(newNode)
		}
	}
}

func (t *TreeNode) Search(value int) *TreeNode {
	if t == nil {
		return nil
	}
	// 1、比较是否为当前节点
	if value == t.Val {
		return t
	}
	// 2、大于当前节点，递归右边
	if value > t.Val {
		return t.Right.Search(value)
	}
	// 3、大于当前节点，递归左边
	if value < t.Val {
		return t.Left.Search(value)
	}
	return nil
}

// Remove 删除一个元素.
func (t *TreeNode) Remove(value int) bool {
	_, existed := remove(t, value)
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

func (t *TreeNode) Min() int {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.Min()
}

func (t *TreeNode) Max() int {
	if t.Right == nil {
		return t.Val
	}
	return t.Right.Max()
}

// 层序遍历
func (t *TreeNode) LevelOrder() [][]int {
	var res [][]int
	if t == nil {
		return res
	}
	q := []*TreeNode{t}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}

func (t *TreeNode) InorderTraversal() []int {
	if t == nil {
		return nil
	}
	var fn func(node *TreeNode)
	var res []int
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Left)
		res = append(res, node.Val)
		fn(node.Right)
	}
	fn(t)
	return res
}
