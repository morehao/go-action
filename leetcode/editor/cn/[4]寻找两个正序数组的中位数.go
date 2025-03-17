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
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	p1, p2 := 0, 0
	sorted := make([]int, 0, m+n)
	for {
		if p1 >= k || p2 >= k {
			break
		}
		// nums1遍历完
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...) // 将 nums2 剩余元素追加到 sorted
			break
		}
		// nums2遍历完
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...) // 将 nums1 剩余元素追加到 sorted
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
	if len(sorted) < k {
		return 0
	}
	return sorted[k-1]
}

// leetcode submit region end(Prohibit modification and deletion)
