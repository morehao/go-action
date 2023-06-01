package main

import (
	"fmt"

	"go-practict/leetcode/linkList"
)

func main() {
	nums := []int{1, 2, 3}
	head := linkList.ArrToLinkList(nums)
	newHead := deleteNode(head, 2)
	fmt.Println(newHead.Scan())
}

func deleteNode(head *linkList.ListNode, val int) *linkList.ListNode {
	if head.Val == val {
		return head.Next
	}
	pre, curr := head, head.Next

	for curr != nil && curr.Val != val {
		pre, curr = curr, curr.Next
	}
	// curr为nil，说明当前链表中不包含对应val；不为nil说明包含。
	if curr != nil {
		pre.Next = curr.Next
	}
	return head
}
