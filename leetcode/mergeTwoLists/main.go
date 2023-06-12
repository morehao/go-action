package main

import (
	"fmt"
	"go-practict/dataStructure/linkList"
)

func main() {
	nums1 := []int{1, 2, 3, 7}
	nums2 := []int{1, 2}
	l1 := linkList.ArrToLinkList(nums1)
	l2 := linkList.ArrToLinkList(nums2)
	fmt.Println(mergeTwoLists(l1, l2).Scan())
}

func mergeTwoLists(l1 *linkList.ListNode, l2 *linkList.ListNode) *linkList.ListNode {
	head := &linkList.ListNode{}
	curr := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return head.Next
}

func mergeTwoLists2(l1 *linkList.ListNode, l2 *linkList.ListNode) *linkList.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}
