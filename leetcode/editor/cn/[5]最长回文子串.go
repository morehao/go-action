/**
给你一个字符串 s，找到 s 中最长的 回文 子串。



 示例 1：


输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。


 示例 2：


输入：s = "cbbd"
输出："bb"




 提示：


 1 <= s.length <= 1000
 s 仅由数字和英文字母组成


 Related Topics 双指针 字符串 动态规划 👍 7583 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0 // 记录最长回文子串的起始和结束位置

	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心，扩展找到最长回文子串
		len1 := expandAroundCenter(s, i, i)   // 奇数长度回文子串
		len2 := expandAroundCenter(s, i, i+1) // 偶数长度回文子串
		maxLen := max(len1, len2)             // 取较长的回文子串长度

		// 如果找到更长的回文子串，更新起始和结束位置
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// 中心扩展函数
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// leetcode submit region end(Prohibit modification and deletion)
