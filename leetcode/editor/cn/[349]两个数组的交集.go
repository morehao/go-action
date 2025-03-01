/**
给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。 

 

 示例 1： 

 
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
 

 示例 2： 

 
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的
 

 

 提示： 

 
 1 <= nums1.length, nums2.length <= 1000 
 0 <= nums1[i], nums2[i] <= 1000 
 

 Related Topics 数组 哈希表 双指针 二分查找 排序 👍 960 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for i := range nums1 {
		m1[nums1[i]] = struct{}{}
	}
	for i := range nums2 {
		m2[nums2[i]] = struct{}{}
	}
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	var res []int
	for k := range m1 {
		if _, ok := m2[k]; ok {
			res = append(res, k)
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
