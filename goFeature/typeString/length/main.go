package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "你好啊"
	fmt.Println(byteLen(s))
	fmt.Println(utf8Len(s))
}

func byteLen(s string) int {
	return len(s)
}

func utf8Len(s string) int {
	return utf8.RuneCountInString(s)
}
