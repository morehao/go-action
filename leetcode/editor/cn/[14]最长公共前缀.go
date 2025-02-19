/**
编写一个函数来查找字符串数组中的最长公共前缀。

 如果不存在公共前缀，返回空字符串 ""。



 示例 1：


输入：strs = ["flower","flow","flight"]
输出："fl"


 示例 2：


输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。



 提示：


 1 <= strs.length <= 200
 0 <= strs[i].length <= 200
 strs[i] 如果非空，则仅由小写英文字母组成


 Related Topics 字典树 字符串 👍 3270 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	prefix := strs[0]
	for i := range strs {
		prefix = getPrefix(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func getPrefix(str1, str2 string) string {
	n := min(len(str1), len(str2))
	i := 0
	for i < n && str1[i] == str2[i] {
		i++
	}
	return str1[:i]
}

// leetcode submit region end(Prohibit modification and deletion)
