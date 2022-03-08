package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(hasSameStr2("abca"))
}

func hasSameStr(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func hasSameStr2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		// ACCII码中127之后的字符，键盘上无法打出
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
