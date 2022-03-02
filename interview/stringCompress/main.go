package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aaabbbbc"
	fmt.Println(compressString(str))
}

func compressString(str string) string {
	if len(str) < 1 {
		return str
	}
	result, cnt := string(str[0]), 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			cnt++
			if i == len(str)-1 {
				result += fmt.Sprintf("%s", strconv.Itoa(cnt))
			}
		} else {
			result += fmt.Sprintf("%s", strconv.Itoa(cnt))
			result += fmt.Sprintf("%s", string(str[i]))
			cnt = 1
			if i == len(str)-1 {
				result += fmt.Sprintf("%s", strconv.Itoa(1))
			}
		}
	}
	if len(result) >= len(str) {
		return str
	}
	return result
}

func compressString2(S string) string {
	var sLen int = len(S)
	if sLen < 2 {
		return S
	}
	var (
		ans []byte = make([]byte, 0, sLen)

		i, j int = 0, 0
	)

	for i < sLen {
		for j < sLen && S[i] == S[j] {
			j++
		}
		ans = append(ans, S[i])
		ans = append(ans, []byte(strconv.Itoa(j-i))...)
		if len(ans) >= sLen {
			return S
		}
		i = j
	}
	return string(ans)
}
