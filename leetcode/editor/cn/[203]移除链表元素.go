/**
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。



 示例 1：


输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]


 示例 2：


输入：head = [], val = 1
输出：[]


 示例 3：


输入：head = [7,7,7,7], val = 7
输出：[]




 提示：


 列表中的节点数目在范围 [0, 10⁴] 内
 1 <= Node.val <= 50
 0 <= val <= 50


 Related Topics 递归 链表 👍 1519 👎 0

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
func removeElements(head *ListNode, val int) *ListNode {
	tempNode := &ListNode{Next: head}
	curr := tempNode
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return tempNode.Next
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)

// leetcode submit region end(Prohibit modification and deletion)
