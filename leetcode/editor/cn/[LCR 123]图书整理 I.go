/**
书店店员有一张链表形式的书单，每个节点代表一本书，节点中的值表示书的编号。为更方便整理书架，店员需要将书单倒过来排列，就可以从最后一本书开始整理，逐一将书放回到
书架上。请倒序返回这个书单链表。 

 

 示例 1： 

 
输入：head = [3,6,4,1]

输出：[1,4,6,3]
 

 

 提示： 

 0 <= 链表长度 <= 10000 

 Related Topics 栈 递归 链表 双指针 👍 484 👎 0

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
func reverseBookList(head *ListNode) []int {
	var res []int
	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
