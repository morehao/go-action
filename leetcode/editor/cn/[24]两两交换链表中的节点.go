/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。 

 

 示例 1： 
 
 
输入：head = [1,2,3,4]
输出：[2,1,4,3]
 

 示例 2： 

 
输入：head = []
输出：[]
 

 示例 3： 

 
输入：head = [1]
输出：[1]
 

 

 提示： 

 
 链表中节点的数目在范围 [0, 100] 内 
 0 <= Node.val <= 100 
 

 Related Topics 递归 链表 👍 2384 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
		Val:  0,
	}

	// prevNode 始终指向待交换的两个节点的前一个节点
	prevNode := dummyHead

	// 当后面至少还有两个节点可以交换时继续循环
	for prevNode.Next != nil && prevNode.Next.Next != nil {
		// 定义三个关键节点指针：
		firstNode := prevNode.Next      // 第一个要交换的节点
		secondNode := prevNode.Next.Next    // 第二个要交换的节点
		nextPairHead := secondNode.Next // 下一对节点的头节点

		// 执行交换操作：
		// 1. 前驱节点指向第二个节点
		prevNode.Next = secondNode
		// 2. 第一个节点指向下一对节点的头节点
		firstNode.Next = nextPairHead
		// 3. 第二个节点指向第一个节点（完成交换）
		secondNode.Next = firstNode

		// 移动prevNode到下一对节点的前一个位置（即当前这对交换后的第二个节点）
		prevNode = firstNode
	}

	// 返回新的头节点（跳过虚拟头节点）
	return dummyHead.Next
}
//leetcode submit region end(Prohibit modification and deletion)
