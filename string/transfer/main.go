package main

import "fmt"

func main() {
	byteToString([]byte{'H', 'e', 'l', 'l', 'o'})
	stringToByte("Hello")
	rangeChineseStr("你好")
	rangeLetterStr("abc")
}

func byteToString(s []byte) string {
	// byte占用1个节字，也就是 8 个比特位，所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符
	result := string(s)
	fmt.Printf("byteToString s:%s\n", result)
	return result
}

func stringToByte(s string) []byte {
	result := []byte(s)
	fmt.Printf("stringToByte s:%s\n", result)
	return result
}

func rangeChineseStr(s string) {
	// rune占用4个字节，共32位比特位，所以它和 int32 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范）
	// 由于 byte 类型能表示的值是有限，只有 2^8=256 个。当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。
	fmt.Printf("rangeChineseStr []rune() len:%d\n", len([]rune(s)))
	fmt.Printf("rangeChineseStr []byte() len:%d\n", len([]byte(s)))
	for i, v := range s {
		fmt.Printf("rangeChineseStr i:%d, v:%s\n", i, string(v))
	}
}

func rangeLetterStr(s string) {
	fmt.Printf("rangeLetterStr []rune() len:%d\n", len([]rune(s)))
	fmt.Printf("rangeLetterStr []byte() len:%d\n", len([]byte(s)))
	for i, v := range []rune(s) {
		fmt.Printf("rangeLetterStr  i:%d, v:%s\n", i, string(v))
	}
}
