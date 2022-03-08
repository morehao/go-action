package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "abc "
	fmt.Println(replaceEmptyStr(s))
}

func replaceEmptyStr(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
