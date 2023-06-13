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
	arr := []int{3, 9, 20, 0, 0, 15, 7}
	tree := BuildTreeWithNums(arr)
	treeByte, _ := jsoniter.Marshal(tree)
	fmt.Println("tree string:", string(treeByte))
	fmt.Println("tree level order:", tree.LevelOrder())
}

func TestBuildTreeWithDataList(t *testing.T) {
	arr := []interface{}{3, 9, 20, nil, nil, 15, 7}
	tree := BuildTreeWithDataList(arr)
	treeByte, _ := jsoniter.Marshal(tree)
	fmt.Println("tree string:", string(treeByte))
	fmt.Println("tree level order:", tree.LevelOrder())
}
