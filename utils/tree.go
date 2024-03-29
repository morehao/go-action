package utils

import (
	"encoding/json"
	"fmt"
)

type NodeItem struct {
	Id     uint64
	Label  string
	Pid    uint64
	Status bool
}

type NodeList []NodeItem

func (l NodeList) ToMap() map[uint64]NodeItem {
	m := make(map[uint64]NodeItem)
	for _, v := range l {
		m[v.Id] = v
	}
	return m
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
		treeMap[node.Id] = &treeItem
		// 根节点收集
		if node.Pid == 0 {
			treeList = append(treeList, &treeItem)
		}
	}
	for _, node := range treeMap {
		if node.Pid != 0 {
			if parent, ok := treeMap[node.Pid]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}
	jsonRes, _ := json.Marshal(treeList)
	fmt.Println(string(jsonRes))
	return treeList
}

func BuildValidTree(nodes []NodeItem) []*NodeTree {
	nodeMap := make(map[uint64]*NodeTree)
	var roots []*NodeTree

	// 创建所有节点
	for i := range nodes {
		nodeMap[nodes[i].Id] = &NodeTree{
			NodeItem: nodes[i],
			Children: make([]*NodeTree, 0),
		}
	}

	// 构建树
	for i := range nodes {
		node := nodeMap[nodes[i].Id]
		if nodes[i].Pid == 0 {
			roots = append(roots, node)
		} else {
			parentNode, ok := nodeMap[nodes[i].Pid]
			if !ok {
				continue
			}
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	// 删除状态为false的节点
	anyRemoved := true
	for anyRemoved {
		anyRemoved = false
		for _, node := range nodeMap {
			for i := 0; i < len(node.Children); {
				child := node.Children[i]
				if !child.Status && len(child.Children) == 0 {
					// 删除无效的叶子节点
					node.Children = append(node.Children[:i], node.Children[i+1:]...)
					anyRemoved = true
				} else {
					i++
				}
			}
		}
	}

	// 过滤出有效的根节点
	validRoots := roots[:0]
	for _, root := range roots {
		if root.Status || len(root.Children) > 0 {
			validRoots = append(validRoots, root)
		}
	}

	jsonRes, _ := json.Marshal(validRoots)
	fmt.Println(string(jsonRes))
	return validRoots
}

func BuildValidTreeWithOption(nodes, validNodes NodeList) []*NodeTree {
	nodeMap := make(map[uint64]*NodeTree)
	var roots []*NodeTree
	validNodeMap := validNodes.ToMap()
	// 创建所有节点
	for i := range nodes {
		nodeMap[nodes[i].Id] = &NodeTree{
			NodeItem: nodes[i],
			Children: make([]*NodeTree, 0),
		}
	}

	// 构建树
	for i := range nodes {
		node := nodeMap[nodes[i].Id]
		if nodes[i].Pid == 0 {
			roots = append(roots, node)
		} else {
			parentNode, ok := nodeMap[nodes[i].Pid]
			if !ok {
				continue
			}
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	// 删除无效的节点
	stack := make([]*NodeTree, len(roots))
	copy(stack, roots)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := 0; i < len(node.Children); i++ {
			_, isValid := validNodeMap[node.Children[i].Id]
			if len(node.Children[i].Children) == 0 && !isValid {
				node.Children = append(node.Children[:i], node.Children[i+1:]...)
				if len(node.Children) > 0 {
					stack = append(stack, node)
				}
				i--
			} else {
				stack = append(stack, node.Children[i])
			}
		}
	}

	jsonRes, _ := json.Marshal(roots)
	fmt.Println(string(jsonRes))
	return roots
}
