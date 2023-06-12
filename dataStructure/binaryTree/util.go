package binaryTree

// 将数组转为二叉树
func ArrayToTree(data []interface{}) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := &TreeNode{}
	switch t := data[0].(type) {
	case int:
		root.Val = t
	case nil:
		return nil
	default:
		panic("Unknown element type")
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	data = data[1:]
	for len(data) > 0 && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 左侧节点
		node.Left = newNodeFromData(data[0])
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		data = data[1:]

		// 右侧节点
		if len(data) > 0 {
			node.Right = newNodeFromData(data[0])
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			data = data[1:]
		}
	}
	return root
}

// 根据数据来创建一个新的节点
func newNodeFromData(val interface{}) *TreeNode {
	switch t := val.(type) {
	case int:
		return &TreeNode{Val: t}
	case nil:
		return nil
	default:
		panic("Unknown element type")
	}
}

func CreateBinaryTree(i int, nums []int) *TreeNode {
	if nums[i] == 0 {
		CreateBinaryTree(i+1, nums)
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

// BuildTree 输入一个切片 ：[3,9,20,0,0,15,7]
func BuildTree(l []int) (root *TreeNode) {
	length := len(l)
	if length == 0 {
		return root
	}

	var nodes = make([]*TreeNode, length)
	root = &TreeNode{
		Val: l[0],
	}
	nodes[0] = root
	//循环输入的数组切片，依次判断每一个节点的左右节点是否存在并创建
	for i := 0; i < length; i++ {
		currentNode := nodes[i]

		if currentNode == nil {
			continue
		}

		leftIndex := 2*i + 1
		if leftIndex < length && l[leftIndex] != 0 {
			currentNode.Left = &TreeNode{
				Val: l[leftIndex],
			}
			nodes[leftIndex] = currentNode.Left
		}

		rightIndex := 2*i + 2
		if rightIndex < length && l[rightIndex] != 0 {
			currentNode.Right = &TreeNode{
				Val: l[rightIndex],
			}
			nodes[rightIndex] = currentNode.Right
		}
	}

	return root
}
