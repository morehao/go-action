package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	arr := []int{5, 3, 6, 2, 4, 0, 0, 1}
	tree := createBinaryTree(0, arr)
	treeStr, _ := jsoniter.Marshal(tree)
	fmt.Println(string(treeStr))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Max())
	fmt.Println(tree.Min())
	fmt.Println(tree.LevelOrder())
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func createBinaryTree(i int, nums []int) *treeNode {
	if nums[i] == 0 {
		return nil
	}
	tree := &treeNode{nums[i], nil, nil}
	// 左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.left = createBinaryTree(2*i+1, nums)
	}
	// 右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.right = createBinaryTree(2*i+2, nums)
	}
	return tree
}

func (nd *treeNode) Insert(newNode *treeNode) {
	if newNode.val == nd.val {
		return
	}
	if newNode.val > nd.val {
		if nd.right == nil {
			nd.right = newNode
		} else {
			nd.right.Insert(newNode)
		}
	} else {
		if nd.left == nil {
			nd.left = newNode
		} else {
			nd.left.Insert(newNode)
		}
	}
}

func (nd *treeNode) Search(value int) *treeNode {
	if nd == nil {
		return nil
	}
	// 1、比较是否为当前节点
	if value == nd.val {
		return nd
	}
	// 2、大于当前节点，递归右边
	if value > nd.val {
		return nd.right.Search(value)
	}
	// 3、大于当前节点，递归左边
	if value < nd.val {
		return nd.left.Search(value)
	}
	return nil
}

// Remove 删除一个元素.
func (nd *treeNode) Remove(value int) bool {
	_, existed := remove(nd, value)
	return existed
}

// 用来递归移除节点的辅助方法.
// 返回替换root的新节点，以及元素是否存在
func remove(root *treeNode, value int) (*treeNode, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	// 从左边找
	if value < root.val {
		root.left, existed = remove(root.left, value)
		return root, existed
	}
	// 从右边找
	if value > root.val {
		root.right, existed = remove(root.right, value)
		return root, existed
	}
	// 如果此节点正是要移除的节点,那么返回此节点，同时返回之前可能需要调整.
	existed = true
	// 如果此节点没有孩子，直接返回即可
	if root.left == nil && root.right == nil {
		root = nil
		return root, existed
	}
	// 如果左子节点为空, 提升右子节点
	if root.left == nil {
		root = root.right
		return root, existed
	}
	// 如果右子节点为空, 提升左子节点
	if root.right == nil {
		root = root.left
		return root, existed
	}
	// 如果左右节点都存在,那么从右边节点找到一个最小的节点提升，这个节点肯定比左子树所有节点都大.
	// 也可以从左子树节点中找一个最大的提升，道理一样.
	smallestInRight := root.right.Min()
	// 提升
	root.val = smallestInRight
	// 从右边子树中移除此节点
	root.right, _ = remove(root.right, smallestInRight)
	return root, existed
}

func (nd *treeNode) Min() int {
	if nd.left == nil {
		return nd.val
	}
	return nd.left.Min()
}

func (nd *treeNode) Max() int {
	if nd.right == nil {
		return nd.val
	}
	return nd.right.Max()
}

// 层序遍历
func (nd *treeNode) LevelOrder() [][]int {
	var ret [][]int
	if nd == nil {
		return ret
	}
	q := []*treeNode{nd}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*treeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.val)
			if node.left != nil {
				p = append(p, node.left)
			}
			if node.right != nil {
				p = append(p, node.right)
			}
		}
		q = p
	}
	return ret
}
