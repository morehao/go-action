/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



 注意：


 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 如果 s 中存在这样的子串，我们保证它是唯一的答案。




 示例 1：


输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。


 示例 2：


输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。


 示例 3:


输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。



 提示：


 m == s.length
 n == t.length
 1 <= m, n <= 10⁵
 s 和 t 由英文字母组成



进阶：你能设计一个在
o(m+n) 时间内解决此问题的算法吗？

 Related Topics 哈希表 字符串 滑动窗口 👍 3195 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// 记录 t 中字符的出现次数
	tCntMap := make(map[byte]int)
	for i := range t {
		tCntMap[t[i]]++
	}

	// 记录窗口内的字符出现次数
	windowCntMap := make(map[byte]int)
	left, right := 0, 0
	start, minLen := 0, len(s)+1
	matchCnt := 0 // 记录窗口中满足条件的字符数量

	for right < len(s) {
		rightChar := s[right] // 右指针字符
		right++

		// 只有 t 中包含该字符，才加入 windowCntMap 进行统计
		if tCntMap[rightChar] > 0 {
			windowCntMap[rightChar]++
			// 只有刚好达到需要数量时，matchCnt 才增加
			if windowCntMap[rightChar] == tCntMap[rightChar] {
				matchCnt++
			}
		}

		// 当窗口内的字符已经满足 t 中所有字符的要求，尝试收缩窗口
		for matchCnt == len(tCntMap) {
			// 记录最小窗口
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			leftChar := s[left] // 左指针字符
			left++

			// 只有 t 中包含该字符，才进行 matchCnt 维护
			if tCntMap[leftChar] > 0 {
				// 只有当移除的字符使窗口匹配字符减少时，matchCnt 才减少
				if windowCntMap[leftChar] == tCntMap[leftChar] {
					matchCnt--
				}
				windowCntMap[leftChar]--
			}
		}
	}

	// 如果 minLen 仍然是初始值，说明没有符合条件的子串
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// leetcode submit region end(Prohibit modification and deletion)
