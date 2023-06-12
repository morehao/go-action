package binaryTree

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestFunc(t *testing.T) {
	arr := []interface{}{1, nil, 2, 3}
	tree := ArrayToTree(arr)
	treeByte, _ := jsoniter.Marshal(tree)
	fmt.Println("tree string:", string(treeByte))
	//fmt.Println("tree search:", tree.Search(4))
	//fmt.Println("tree max:", tree.Max())
	//fmt.Println("tree min:", tree.Min())
	//fmt.Println("tree level order:", tree.LevelOrder())
}
