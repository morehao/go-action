package transfer

import (
	"encoding/json"
	"fmt"
)

type NodeItem struct {
	Id    uint64
	Label string
	Pid   uint64
}

type NodeTree struct {
	NodeItem
	Children []*NodeTree
}

// type NodeTreeMap map[uint64]*NodeTree

func BuildTree(nodes []NodeItem) []*NodeTree {
	treeList := make([]*NodeTree, 0)
	treeMap := make(map[uint64]*NodeTree)
	for _, node := range nodes {
		treeItem := NodeTree{
			NodeItem: NodeItem{
				Id:    node.Id,
				Pid:   node.Pid,
				Label: node.Label,
			},
			Children: make([]*NodeTree, 0),
		}
		// 根节点收集
		if node.Pid == 0 {
			treeList = append(treeList, &treeItem)
		} else {
			// 子节点收集
			if _, ok := treeMap[node.Pid]; ok {
				treeMap[node.Pid].Children = append(treeMap[node.Pid].Children, &treeItem)
			}
		}
		// 把节点映射到map表
		treeMap[node.Id] = &treeItem
	}
	jsonRes, _ := json.Marshal(treeList)
	fmt.Println(string(jsonRes))
	return treeList
}
