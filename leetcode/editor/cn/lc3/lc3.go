/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。



 示例 1:


输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


 示例 2:


输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


 示例 3:


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。




 提示：


 0 <= s.length <= 5 * 10⁴
 s 由英文字母、数字、符号和空格组成


 Related Topics 哈希表 字符串 滑动窗口 👍 10492 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left := 0
	m := make(map[byte]int) // 使用 map 存储字符的最后出现位置
	for right := range s {
		rightChar := s[right]
		// 如果当前字符曾经出现过，更新 left 指针的位置
		if lastPos, ok := m[rightChar]; ok && lastPos >= left {
			left = lastPos + 1
		}
		m[rightChar] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

// leetcode submit region end(Prohibit modification and deletion)
