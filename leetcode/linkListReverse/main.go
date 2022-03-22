package main

import (
	"fmt"

	"go-practict/leetcode/linkList"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	linkNode := linkList.ArrToLinkList(nums)
	fmt.Println(linkNode.Scan())
	newLinkNode := reverseList(linkNode)
	fmt.Println(newLinkNode.Scan())
}

func reverseList(head *linkList.ListNode) *linkList.ListNode {
	var pre *linkList.ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}
