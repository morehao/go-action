/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

 算法的时间复杂度应该为 O(log (m+n)) 。



 示例 1：


输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2


 示例 2：


输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5






 提示：


 nums1.length == m
 nums2.length == n
 0 <= m <= 1000
 0 <= n <= 1000
 1 <= m + n <= 2000
 -10⁶ <= nums1[i], nums2[i] <= 10⁶


 Related Topics 数组 二分查找 分治 👍 7442 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := mergeNums(nums1, nums2)
	n := len(nums)
	if n%2 == 1 {
		return float64(nums[n/2])
	} else {
		return float64(nums[n/2]+nums[n/2-1]) / 2
	}
}

func mergeNums(nums1, nums2 []int) []int {
	p1, p2 := 0, 0
	m, n := len(nums1), len(nums2)
	var sorted []int
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	return sorted
}

// leetcode submit region end(Prohibit modification and deletion)
