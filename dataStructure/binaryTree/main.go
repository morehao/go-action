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
