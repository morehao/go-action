/**
给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。



 示例 1：


输入：head = [2,4,7,8], cnt = 1
输出：8



 提示：


 1 <= head.length <= 100
 0 <= head[i] <= 100
 1 <= cnt <= head.length




 Related Topics 链表 双指针 👍 527 👎 0

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
func trainingPlan(head *ListNode, cnt int) *ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	kthNode := head
	// 倒数第 cnt 个节点相当于是第 n-cnt 个节点
	for n > cnt {
		kthNode = kthNode.Next
		n--
	}
	return kthNode
}

// leetcode submit region end(Prohibit modification and deletion)
