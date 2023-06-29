package main

var cacheNodeMap = map[*Node]*Node{}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cacheNodeMap[node]; ok {
		return n
	}
	newNode := &Node{
		Val: node.Val,
	}
	cacheNodeMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

// 回溯 + 哈希表
func copyRandomList1(head *Node) *Node {
	return deepCopy(head)
}
