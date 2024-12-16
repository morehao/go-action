package main

import (
	"fmt"
	"go-practict/dataStructure/binaryTree"
	"strconv"
	"strings"
)

func main() {
	list := []interface{}{1, 2, 3, nil, nil, 4, 5}
	root := binaryTree.BuildTreeWithArray(list)
	fmt.Println(root.LevelOrder())
	codec := Constructor()
	serializeStr := codec.serialize(root)
	fmt.Println(serializeStr)
	deserializeNode := codec.deserialize(serializeStr)
	fmt.Println(deserializeNode.LevelOrder())
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *binaryTree.TreeNode) string {
	var (
		res strings.Builder
		dfs func(node *binaryTree.TreeNode)
	)
	dfs = func(node *binaryTree.TreeNode) {
		if node == nil {
			res.WriteString("null,")
			return
		}
		res.WriteString(strconv.Itoa(node.Val))
		res.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *binaryTree.TreeNode {
	var (
		list      = strings.Split(data, ",")
		builderFn func() *binaryTree.TreeNode
	)
	builderFn = func() *binaryTree.TreeNode {
		if len(list) == 0 {
			return nil
		}
		if list[0] == "null" {
			list = list[1:]
			return nil
		}
		v, _ := strconv.Atoi(list[0])
		list = list[1:]
		return &binaryTree.TreeNode{
			Val:   v,
			Left:  builderFn(),
			Right: builderFn(),
		}
	}
	return builderFn()
}
