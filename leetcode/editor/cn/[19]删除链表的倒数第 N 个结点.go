/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



 示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]


 示例 2：


输入：head = [1], n = 1
输出：[]


 示例 3：


输入：head = [1,2], n = 1
输出：[1]




 提示：


 链表中结点的数目为 sz
 1 <= sz <= 30
 0 <= Node.val <= 100
 1 <= n <= sz




 进阶：你能尝试使用一趟扫描实现吗？

 Related Topics 链表 双指针 👍 3073 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dummyHead := &ListNode{
		Next: head,
		Val:  0,
	}
	curr := dummyHead
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	pre := nodes[len(nodes)-n-1]
	pre.Next = pre.Next.Next
	return dummyHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
