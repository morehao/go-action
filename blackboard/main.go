package main

import (
	"fmt"
	"unsafe"

	"go-practict/leetcode/linkList"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	linkNode := linkList.ArrToLinkList(nums)
	fmt.Println(linkNode.Scan())
	newLinkNode := reverseList(linkNode)
	fmt.Println(newLinkNode.Scan())
}

type slice struct {
	array []unsafe.Pointer
	len   int
	cap   int
}

type stringStruct struct {
	str unsafe.Pointer
	len int
}

type hmap struct {
	count      int            // 当前保存的元素个数
	B          uint8          // bucket数组的大小
	buckets    unsafe.Pointer // bucket数组，数组的长度为2的B次方
	oldbuckets unsafe.Pointer // 老旧的buckets数组，用于扩容
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
