package main

import (
	"fmt"

	"go-practict/leetcode/linkList"
)

func main() {
	nums := []int{1, 2, 3, 7}
	head := linkList.ArrToLinkList(nums)
	fmt.Println(getKthFromEnd(head, 1).Scan())
}

/*
1、首先将fast指向链表的头节点，然后向后走k步，则此时fast指针刚好指向链表的第k+1个节点。
2、我们首先将slow指向链表的头节点，同时slow 与fast 同步向后走，当fast 指针指向链表的尾部空节点时，则此时返回slow所指向的节点即可。
*/
func getKthFromEnd(head *linkList.ListNode, k int) *linkList.ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	// fast为空表示链表已遍历结束
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

/*
1、获取链表长度；
2、计算出倒数第k个节点，即下标为n-k的节点。
*/
func getKthFromEnd1(head *linkList.ListNode, k int) *linkList.ListNode {
	len := 0
	for node := head; node != nil; node = node.Next {
		len++
	}
	var target *linkList.ListNode
	target = head
	for len > k {
		target = target.Next
		len--
	}
	return target
}
