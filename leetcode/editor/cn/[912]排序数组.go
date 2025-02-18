/**
给你一个整数数组 nums，请你将该数组升序排列。

 你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。






 示例 1：


输入：nums = [5,2,3,1]
输出：[1,2,3,5]


 示例 2：


输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]




 提示：


 1 <= nums.length <= 5 * 10⁴
 -5 * 10⁴ <= nums[i] <= 5 * 10⁴


 Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 1077 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	partition(nums, 0, len(nums)-1)
	return nums
}

/*
快速排序使用分治法策略来把一个序列分为较小和较大的 2 个子序列，然后递回地排序两个子序列。具体算法描述如下：
1、从序列中随机挑出一个元素，做为 “基准”(pivot)；
2、重新排列序列，将所有比基准值小的元素摆放在基准前面，所有比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个操作结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3、递归地把小于基准值元素的子序列和大于基准值元素的子序列进行快速排序。
*/
func partition(nums []int, start, end int) {
	i, j := start, end
	midValue := nums[(start+end)/2]
	for i <= j {
		for nums[i] < midValue {
			i++
		}
		for nums[j] > midValue {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if i < end {
		partition(nums, i, end)
	}
	if j > start {
		partition(nums, start, j)
	}
}

// leetcode submit region end(Prohibit modification and deletion)
