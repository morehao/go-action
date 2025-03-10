/**
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

 如果数组中不存在目标值 target，返回 [-1, -1]。

 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



 示例 1：


输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

 示例 2：


输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

 示例 3：


输入：nums = [], target = 0
输出：[-1,-1]



 提示：


 0 <= nums.length <= 10⁵
 -10⁹ <= nums[i] <= 10⁹
 nums 是一个非递减数组
 -10⁹ <= target <= 10⁹


 Related Topics 数组 二分查找 👍 2932 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	left := fn(nums, target)    // 找到 target 的最左索引
	right := fn(nums, target+1) // 找到比 target 大的第一个元素索引
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1} // 说明 target 不存在
	}
	return []int{left, right - 1} // 右边界修正为 right-1
}

// 查找 >= target 的最左索引
func fn(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left // 返回 >= target 的第一个位置
}

// leetcode submit region end(Prohibit modification and deletion)
