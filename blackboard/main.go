package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	rk, res := -1, 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < len(s) && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		res = max(res, rk-i+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
