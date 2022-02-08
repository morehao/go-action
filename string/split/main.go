package main

import "fmt"

func main() {
	splitChineseStr("你好世界")
	splitLetterStr("abcd")
}

func splitChineseStr(s string) {
	// 中文字符串截取，需要使用[]rune(s)转为unicode码，然后切割
	fmt.Printf("splitChineseStr s:%s\n", string([]rune(s)[:2]))
}

func splitLetterStr(s string) {
	// 英文字符串可以直接切割也可以转为unicode码之后再切割
	fmt.Printf("splitLetterStr s1:%s\n", s[:2])
	fmt.Printf("splitLetterStr s2:%s\n", string([]rune(s)[:2]))
}
