package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 忽略大小写
	s = strings.ToUpper(s)
	var newBytes []byte
	for _, v := range []byte(s) {
		if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z') {
			newBytes = append(newBytes, v)
		}
	}
	left, right := 0, len(newBytes)-1
	for left <= right {
		// 回文字符串左右对称
		if newBytes[left] == newBytes[right] {
			left++
			right--
			continue
		} else {
			return false
		}
	}
	return true
}
