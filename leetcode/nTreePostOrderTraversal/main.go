package main

func postorder(root *Node) []int {
	res := make([]int, 0)
	var order func(node *Node)
	order = func(node *Node) {
		if node == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			order(node.Children[i])
		}
		res = append(res, node.Val)
	}
	order(root)
	return res
}

type Node struct {
	Val      int
	Children []*Node
}
