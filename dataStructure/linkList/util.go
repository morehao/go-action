package linkList

func ArrToLinkList(nums []int) *ListNode {
	if len(nums) < 1 {
		return nil
	}
	var header *ListNode
	current := header
	for i := 0; i < len(nums); i++ {
		node := &ListNode{
			Val: nums[i],
		}
		if header == nil {
			header = node
			current = header
		} else {
			current.Next = node
			current = current.Next
		}
	}
	return header
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}
