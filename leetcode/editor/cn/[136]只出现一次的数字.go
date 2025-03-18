/**
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。







 示例 1 ：


 输入：nums = [2,2,1]


 输出：1

 示例 2 ：


 输入：nums = [4,1,2,1,2]


 输出：4

 示例 3 ：


 输入：nums = [1]


 输出：1



 提示：


 1 <= nums.length <= 3 * 10⁴
 -3 * 10⁴ <= nums[i] <= 3 * 10⁴
 除了某个元素只出现一次以外，其余每个元素均出现两次。


 Related Topics 位运算 数组 👍 3284 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	// 异或性质：
	// x ^ x = 0（相同数异或为 0）
	// x ^ 0 = x（任何数与 0 异或不变）
	// 异或满足交换律和结合律，所以 a ^ b ^ a 计算顺序不影响结果
	var res int
	for i := range nums {
		res = res ^ nums[i]
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
