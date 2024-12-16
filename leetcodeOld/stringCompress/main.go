package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aabcccccaaa"
	fmt.Println(compressString(str))
}

func compressString(S string) string {
	res := make([]byte, 0)
	i, j := 0, 0
	size := len(S)
	for i < size {
		for j < size && S[j] == S[i] {
			j++
		}
		res = append(res, S[i])
		res = append(res, []byte(strconv.Itoa(j-i))...)
		i = j
	}
	if len(res) >= size {
		return S
	}
	return string(res)
}
