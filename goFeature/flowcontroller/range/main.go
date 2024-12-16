package main

import "fmt"

func main() {
	rangEnglishStr()
	rangeChineseStr()
}

func rangEnglishStr() {
	s := "abc"
	for i, v := range s {
		fmt.Printf("rangEnglishStr index:%d, value:%c\n", i, v)
	}
}

func rangeChineseStr() {
	s := "你好"
	for i, v := range s {
		fmt.Printf("rangeChineseStr index:%d, value:%c\n", i, v)
	}
}
