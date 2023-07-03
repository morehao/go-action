package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstringWithSlideWindow(s))
}

/*
1、状态定义：设动态规划列表dp,dp[i]代表以字符s为结尾的“最长不重复子字符串”的长度。
2、转移方程：固定右边界j,设字符s[j]左边距离最近的相同字符为s[i],即s[j]=s[i]。
	1.当s[i]不在map中,即s[j]左边无相同字符，则dp[j]=dp[j-1]+1；
	2.当dp[j-1]<j-i,说明字符s[i]在子字符串dp[i-1]区间之外，则dp[j]=dp[j-1]+1；
	3.当dp[j-1]≥j-i,说明字符s[i]在子字符串dp[j-i]区间之中，则dp[j]的左边界由s[i]决定，即dp[j]=j-i；
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	rk := -1
	res := 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < len(s) && !m[s[rk+1]] {
			m[s[rk+1]] = true
			rk++
			if rk+1-i > res {
				res = rk + 1 - i
			}
		}
	}
	return res
}

// 滑动窗口
func lengthOfLongestSubstringWithSlideWindow(s string) int {
	n := len(s)
	// 哈希集合，记录每个字符是否出现过
	m := make(map[byte]bool)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, res := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && !m[s[rk+1]] {
			m[s[rk+1]] = true
			rk++
			// 第 i 到 rk 个字符是一个极长的无重复字符子串
			res = max(res, rk-i+1)
		}
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
