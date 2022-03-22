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
	res := make([]byte, 0, len(str))
	i, j := 0, 0
	sLen := len(str)
	for i < sLen {
		for j < sLen && str[j] == str[i] {
			j++
		}
		res = append(res, str[i])
		res = append(res, []byte(strconv.Itoa(j-i))...)
		if len(res) >= sLen {
			return str
		}
		i = j
	}
	return string(res)
}
