package main

import (
	"fmt"
	"go-practict/dataStructure/linkList"
)

func main() {
	nums := []int{1, 2, 3}
	head := linkList.ArrToLinkList(nums)
	fmt.Println(reversePrint(head))
}

func reversePrint(head *linkList.ListNode) []int {
	res := make([]int, 0)
	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
