/**
教练使用整数数组 actions 记录一系列核心肌群训练项目编号。为增强训练趣味性，需要将所有奇数编号训练项目调整至偶数编号训练项目之前。请将调整后的训练项目编
号以 数组 形式返回。



 示例 1：


输入：actions = [1,2,3,4,5]
输出：[1,3,5,2,4]
解释：为正确答案之一



 提示：


 0 <= actions.length <= 50000
 0 <= actions[i] <= 10000




 Related Topics 数组 双指针 排序 👍 339 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
// 双指针
func trainingPlan(actions []int) []int {
	i := 0 // 标记当前已经调整好的奇数编号训练项目的末尾位置
	for j := 0; j < len(actions); j++ {
		if actions[j]%2 == 1 {
			actions[j], actions[i] = actions[i], actions[j]
			i++
		}
	}
	return actions
}

// leetcode submit region end(Prohibit modification and deletion)
