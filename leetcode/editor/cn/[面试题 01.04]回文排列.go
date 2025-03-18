/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

 回文串不一定是字典当中的单词。



 示例1：

 输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）




 Related Topics 位运算 哈希表 字符串 👍 138 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	// 统计出现次数为奇数的字符数量
	oddCount := 0
	for _, count := range m {
		if count%2 != 0 {
			oddCount++
		}
	}
	if len(s)%2 == 0 {
		// 偶数长度，所有字符出现次数必须为偶数
		return oddCount == 0
	}
	// 奇数长度，只有一个字符出现次数为奇数
	return oddCount == 1
}

// leetcode submit region end(Prohibit modification and deletion)
