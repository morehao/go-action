package linkList

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func (l *ListNode) Scan() []int {
	res := make([]int, 0)
	current := l
	i := 1
	res = append(res, current.Val)
	for current.Next != nil {
		current = current.Next
		res = append(res, current.Val)
		i++
	}
	return res
}
