/**
URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使
用字符数组实现，以便直接在数组上操作。）



 示例 1：


输入："Mr John Smith    ", 13
输出："Mr%20John%20Smith"


 示例 2：


输入："               ", 5
输出："%20%20%20%20%20"




 提示：


 字符串长度在 [0, 500000] 范围内。


 Related Topics 字符串 👍 123 👎 0

*/

package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func replaceSpaces(S string, length int) string {
	var res strings.Builder
	for i := 0; i < length; i++ {
		char := S[i]
		if char != ' ' {
			res.WriteByte(char)
		} else {
			res.WriteString("%20")
		}
	}
	return res.String()
}

// leetcode submit region end(Prohibit modification and deletion)
