package main

import "fmt"

func main() {
	s := "abcdefg"
	fmt.Println(reverseLeftWords(s, 2))
}

func reverseLeftWords(s string, n int) string {
	if len(s) <= n {
		return s
	}
	res := make([]byte, 0, len(s))
	str := []byte(s)
	res = append(res, str[n:]...)
	res = append(res, str[:n]...)
	return string(res)
}
