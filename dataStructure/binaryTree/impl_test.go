package binaryTree

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestFunc(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinaryTree(0, arr)
	treeByte, _ := jsoniter.Marshal(tree)
	fmt.Println("tree string:", string(treeByte))
	fmt.Println("tree search:", tree.Search(4))
	fmt.Println("tree max:", tree.Max())
	fmt.Println("tree min:", tree.Min())
	fmt.Println("tree level order:", tree.LevelOrder())
}
