/**
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。 

 本题中，将空字符串定义为有效的 回文串 。 

 

 示例 1： 

 
输入: s = "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串 

 示例 2： 

 
输入: s = "race a car"
输出: false
解释："raceacar" 不是回文串 

 

 提示： 

 
 1 <= s.length <= 2 * 10⁵ 
 字符串 s 由 ASCII 字符组成 
 

 

 
 注意：本题与主站 125 题相同： https://leetcode-cn.com/problems/valid-palindrome/ 

 Related Topics 双指针 字符串 👍 71 👎 0

*/

package main

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	// 取出所有字母
	var newBytes []byte
	s = strings.ToUpper(s)
	for i := range s {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			newBytes = append(newBytes, s[i])
		}
	}
	left, right := 0, len(newBytes)-1
	for left <= right {
		if newBytes[left] == newBytes[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
