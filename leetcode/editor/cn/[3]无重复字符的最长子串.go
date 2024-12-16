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
func lengthOfLongestSubstring(s string) int {
	// 窗口集合，用来存储当前窗口中的字符
	windowsMap := make(map[byte]struct{})
	left := 0
	maxLength := 0

	// 遍历字符串，right 是右边指针
	for right := 0; right < len(s); right++ {
		rightChar := s[right]
		// 如果当前字符在窗口中已经出现，移动左指针
		for _, exists := windowsMap[rightChar]; exists; {
			delete(windowsMap, s[left])       // 删除左边字符
			left++                            // 左指针向右移动
			_, exists = windowsMap[rightChar] // 检查新的字符
		}

		// 将当前字符加入窗口
		windowsMap[rightChar] = struct{}{}
		// 更新最大子串长度
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

// leetcode submit region end(Prohibit modification and deletion)
