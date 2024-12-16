package binaryTree

// BuildTreeWithArray
// @Description: 通过interface数组构建二叉树
// @param data interface数组
// @return *TreeNode 构建好的二叉树
func BuildTreeWithArray(data []interface{}) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	root := newNodeFromData(data[0])
	queue := []*TreeNode{root}
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

// CreateBinaryTree
// @Description: 通过数组创建一个完全二叉树
// @param i 数组中元素下标
// @param nums 数组
// @return *TreeNode 完全二叉树节点
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

// BuildTreeWithNums 输入一个切片 ：[3,9,20,0,0,15,7]
func BuildTreeWithNums(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[0],
	}
	nodes := make([]*TreeNode, length)
	nodes[0] = root
	// 循环输入的数组切片，依次判断每一个节点的左右节点是否存在并创建
	for i := 0; i < length; i++ {
		currentNode := nodes[i]
		if currentNode == nil {
			continue
		}
		var (
			leftIndex  = 2*i + 1
			rightIndex = 2*i + 2
		)
		if leftIndex < length && nums[leftIndex] != 0 {
			currentNode.Left = &TreeNode{
				Val: nums[leftIndex],
			}
			nodes[leftIndex] = currentNode.Left
		}
		if rightIndex < length && nums[rightIndex] != 0 {
			currentNode.Right = &TreeNode{
				Val: nums[rightIndex],
			}
			nodes[rightIndex] = currentNode.Right
		}
	}
	return root
}

func BuildTreeWithDataList(dataList []interface{}) *TreeNode {
	length := len(dataList)
	if length == 0 {
		return nil
	}
	root := newNodeFromData(dataList[0])
	nodes := make([]*TreeNode, length)
	nodes[0] = root
	// 循环输入的数组切片，依次判断每一个节点的左右节点是否存在并创建
	for i := 0; i < length; i++ {
		currentNode := nodes[i]
		if currentNode == nil {
			continue
		}
		var (
			leftIndex  = 2*i + 1
			rightIndex = 2*i + 2
		)
		if leftIndex < length && dataList[leftIndex] != nil {
			currentNode.Left = newNodeFromData(dataList[leftIndex])
			nodes[leftIndex] = currentNode.Left
		}
		if rightIndex < length && dataList[rightIndex] != nil {
			currentNode.Right = newNodeFromData(dataList[rightIndex])
			nodes[rightIndex] = currentNode.Right
		}
	}
	return root
}
