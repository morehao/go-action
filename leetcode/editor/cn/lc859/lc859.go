/**
给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。

 交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。


 例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。




 示例 1：


输入：s = "ab", goal = "ba"
输出：true
解释：你可以交换 s[0] = 'a' 和 s[1] = 'b' 生成 "ba"，此时 s 和 goal 相等。

 示例 2：


输入：s = "ab", goal = "ab"
输出：false
解释：你只能交换 s[0] = 'a' 和 s[1] = 'b' 生成 "ba"，此时 s 和 goal 不相等。

 示例 3：


输入：s = "aa", goal = "aa"
输出：true
解释：你可以交换 s[0] = 'a' 和 s[1] = 'a' 生成 "aa"，此时 s 和 goal 相等。




 提示：


 1 <= s.length, goal.length <= 2 * 10⁴
 s 和 goal 由小写英文字母组成


 Related Topics 哈希表 字符串 👍 322 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		// 如果相同，需要有重复字符串，才能确保交换后和 goal 相同
		m := make(map[byte]struct{})
		for i := range s {
			if _, ok := m[s[i]]; ok {
				return true
			}
			m[s[i]] = struct{}{}
		}
		return false
	}
	var diff []int
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}
	if len(diff) != 2 {
		return false
	}
	// 检查交换这两个位置的字符后是否相等
	return s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}

// leetcode submit region end(Prohibit modification and deletion)
