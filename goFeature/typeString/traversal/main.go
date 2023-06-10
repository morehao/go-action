package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "hello 你好"
	for _, v := range s {
		ctype := reflect.TypeOf(v)
		// rune类型||int32类型
		// go的源代码使用UTF-8编码，所以字符串中的文字的源代码也是UTF-8文本。
		fmt.Printf("for range type:%s\n", ctype)
	}
	for i := 0; i < len(s); i++ {
		ctype := reflect.TypeOf(s[i])
		// byte类型||uint8类型
		fmt.Printf("for index type:%s\n", ctype)
	}
}
