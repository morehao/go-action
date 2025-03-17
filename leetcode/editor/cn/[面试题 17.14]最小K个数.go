/**
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

 示例：

 输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]


 提示：


 0 <= len(arr) <= 100000
 0 <= k <= min(100000, len(arr))


 Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 241 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k > len(arr) {
		return nil
	}
	partition(arr, 0, len(arr)-1)
	return arr[:k]
}

func partition(nums []int, start, end int) {
	i, j := start, end
	mid := (start + end) / 2
	midVal := nums[mid]
	for i <= j {
		for nums[i] < midVal {
			i++
		}
		for nums[j] > midVal {
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
