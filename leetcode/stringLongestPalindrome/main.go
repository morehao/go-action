package main

import "fmt"

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}

/*
1.回文字符串长度为偶数时，类似这样的字符串abba;长度为奇数时，有个中心字符，类似abcba，c为中心字符;
2.如果一个字符出现次数x为偶数，那可以使回文字符串长度加x；
如果出现次数x是奇数，不考虑中心字符，可以使回文字符串长度加x-1，再加上考虑中心字符，则回文字符串长度可以再加1；
*/
func longestPalindrome(s string) int {
	m := make(map[byte]int)
	// 是否有奇数次数标识
	var flag bool
	var ans int
	for i := range s {
		m[s[i]]++
	}
	for _, v := range m {
		if v%2 != 0 {
			flag = true
			ans += v - 1
			continue
		}
		ans += v
	}
	if flag {
		ans++
	}
	return ans
}
