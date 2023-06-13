package binaryTree

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	arr := []int{1, 2, 3}
	tree := CreateBinaryTree(0, arr)
	treeByte, _ := jsoniter.Marshal(tree)
	fmt.Println("tree string:", string(treeByte))
	fmt.Println("tree level order:", tree.LevelOrder())
}

func TestBuildTree(t *testing.T) {
	arr := []int{1, 0, 2, 3}
	tree := BuildTree(arr)
	treeByte, _ := jsoniter.Marshal(tree)
	fmt.Println("tree string:", string(treeByte))
	fmt.Println("tree level order:", tree.LevelOrder())
}
