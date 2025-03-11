/**
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。


 返回 滑动窗口中的最大值 。



 示例 1：


输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


 示例 2：


输入：nums = [1], k = 1
输出：[1]




 提示：


 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴
 1 <= k <= nums.length


 Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 3045 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	// 结果数组
	result := make([]int, 0, len(nums)-k+1)
	// 双端队列，存储数组元素的索引
	deque := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// 移除队列中不在当前窗口的元素，窗口的范围为 [i-k+1, i]
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// 移除队列中所有小于当前元素的元素，保持队列单调递减
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 将当前元素的索引加入队列
		deque = append(deque, i)

		// 当窗口大小达到 k 时，将队列头部的元素加入结果数组
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)
