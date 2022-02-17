package main

import "fmt"

func main() {
	s := "string"
	// 很多函数的参数为interface类型，比如fmt.Println(a ...interface{}),编译期间很难确定其参数的具体类型，也会发生逃逸
	fmt.Println(s)
}
