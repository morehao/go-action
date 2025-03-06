/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变
短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

 示例 1：


输入："aabcccccaaa"
输出："a2b1c5a3"


 示例 2：


输入："abbccd"
输出："abbccd"
解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。


 提示：


 字符串长度在 [0, 50000] 范围内。


 Related Topics 双指针 字符串 👍 189 👎 0

*/

package main

import "strconv"

// leetcode submit region begin(Prohibit modification and deletion)
func compressString(S string) string {
	n := len(S)
	if n == 0 {
		return S
	}
	var res []byte
	count := 1
	for i := 1; i < n; i++ {
		if S[i] == S[i-1] {
			count++
		} else {
			res = append(res, S[i-1])
			res = append(res, []byte(strconv.Itoa(count))...)
			count = 1
		}
	}
	res = append(res, S[n-1])
	res = append(res, []byte(strconv.Itoa(count))...)
	if len(res) >= n {
		return S
	}
	return string(res)

}

// leetcode submit region end(Prohibit modification and deletion)
