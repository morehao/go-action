/**
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

 示例 1：


 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True


 示例 2：


 输入：s1 = "aa", s2 = "aba"
 输出：False





 提示：


 字符串长度在[0, 100000]范围内。


 说明:


 你能只调用一次检查子串的方法吗？


 Related Topics 字符串 字符串匹配 👍 257 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func isFlipedString(s1 string, s2 string) bool {
	// 如果长度不相等，s2 不可能是 s1 的旋转
	m, n := len(s1), len(s2)
	if m != n {
		return false
	}

	// 将 s1 与自身拼接
	s := s1 + s1

	// 手动检查 s2 是否是 s 的子串
	for i := 0; i+n <= len(s); i++ {
		if s[i:i+n] == s2 {
			return true
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)
