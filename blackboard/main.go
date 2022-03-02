package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aaabbbbc"
	fmt.Println(compressString(str))
}

func compressString(s string) string {
	sLen := len(s)
	if len(s) < 2 {
		return s
	}
	var (
		i, j = 0, 0
		ans  = make([]byte, 0, sLen)
	)
	for i < sLen {
		for j < sLen && s[i] == s[j] {
			j++
		}
		ans = append(ans, s[i])
		ans = append(ans, []byte(strconv.Itoa(j-i))...)
		i = j
	}
	if len(ans) >= sLen {
		return s
	}
	return string(ans)
}
