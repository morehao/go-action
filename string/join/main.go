package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string
	s = "中国"
	var b bytes.Buffer
	b.WriteString("中国")
	for i := 0; i < 10; i++ {
		// 原生方式：string不可变，string按增量方式构建字符串会导致多次内存分配和复制
		s += "a"
		// bytes.Buffer是一个变长的字节缓存区。其内部使用slice来存储字节，
		// 使用WriteString进行字符串拼接时，其会根据情况动态扩展slice长度，相较于原生方式，更加高效
		b.WriteString("a")
	}
	fmt.Println(s)
	fmt.Println(b.String())
}
