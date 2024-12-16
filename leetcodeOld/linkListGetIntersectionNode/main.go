package main

import "go-practict/dataStructure/linkList"

/*
双指针法
链表
headA和headB的长度分别是m和n。假设链表headA的不相交部分有a个节点，链表headB的不相交部分有b个节点，两个链表相交的部分有c个节点，则有a+c=m，b+c=n。
如果p1遍历完指向headB，相当于移动了a+c+b次，同理，p2遍历完相当于移动了b+c+a次，遍历到末尾就可以得到相交节点或nil。
*/
func getIntersectionNode(headA, headB *linkList.ListNode) *linkList.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}

func getIntersectionNode1(headA, headB *linkList.ListNode) *linkList.ListNode {
	m := make(map[*linkList.ListNode]struct{})
	for node := headA; node != nil; node = node.Next {
		m[node] = struct{}{}
	}
	for node := headB; node != nil; node = node.Next {
		if _, ok := m[node]; ok {
			return node
		}
	}
	return nil
}
