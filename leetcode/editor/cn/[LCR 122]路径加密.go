/**
假定一段路径记作字符串 path，其中以 "." 作为分隔符。现需将路径加密，加密方法为将 path 中的分隔符替换为空格 " "，请返回加密后的字符串。



 示例 1：


输入：path = "a.aef.qerf.bb"

输出："a aef qerf bb"





 限制：

 0 <= path.length <= 10000

 Related Topics 字符串 👍 584 👎 0

*/

package main

// leetcode submit region begin(Prohibit modification and deletion)
func pathEncryption(path string) string {
	bytes := []byte(path)
	for i := range bytes {
		if bytes[i] == '.' {
			bytes[i] = ' '
		}
	}
	return string(bytes)
}

// leetcode submit region end(Prohibit modification and deletion)
