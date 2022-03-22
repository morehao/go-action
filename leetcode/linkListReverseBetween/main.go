package main

import (
	"fmt"

	"go-practict/leetcode/linkList"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	linkNode := linkList.ArrToLinkList(nums)
	fmt.Println(linkNode.Scan())
	newLinkNode := reverseListBetween(linkNode, 2, 4)
	fmt.Println(newLinkNode.Scan())
}

func reverseListBetween(head *linkList.ListNode, left, right int) *linkList.ListNode {
	dummyNode := &linkList.ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
