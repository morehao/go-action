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
	var res []int
	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
