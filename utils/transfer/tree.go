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

func BuildTree() []*NodeTree {
	var items = []NodeItem{
		{Id: 1, Pid: 0, Label: "1"},
		{Id: 2, Pid: 1, Label: "1-2"},
		{Id: 3, Pid: 1, Label: "1-3"},
		{Id: 4, Pid: 3, Label: "1-3-4"},
		{Id: 5, Pid: 0, Label: "5"},
		{Id: 6, Pid: 5, Label: "5-6"},
		{Id: 7, Pid: 6, Label: "5-6-7"},
		{Id: 8, Pid: 6, Label: "5-6-8"},
	}
	treeList := make([]*NodeTree, 0)
	treeMap := make(map[uint64]*NodeTree)
	for _, item := range items {
		treeItem := NodeTree{
			NodeItem: NodeItem{
				Id:    item.Id,
				Pid:   item.Pid,
				Label: item.Label,
			},
			Children: make([]*NodeTree, 0),
		}
		// 根节点收集
		if item.Pid == 0 {
			treeList = append(treeList, &treeItem)
		} else {
			if item.Pid == 2 {
				fmt.Println(2)
			}
			// 子节点收集
			treeMap[item.Pid].Children = append(treeMap[item.Pid].Children, &treeItem)

		}
		// 把节点映射到map表
		treeMap[item.Id] = &treeItem
	}
	jsonRes, _ := json.Marshal(treeList)
	fmt.Println(string(jsonRes))
	return treeList
}
