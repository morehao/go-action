package main

import "fmt"

func main() {
	// 先声明再赋值
	var s1 string
	s1 = "Hello World"
	fmt.Printf("s1:%s\n", s1)
	//	简短变量声明
	s2 := "Hello World"
	fmt.Printf("s2:%s\n", s2)

	//	双引号声明
	s3 := "Hi, \nTim"
	fmt.Printf("s3:%s\n", s3)
	s4 := `Hi,
Tim`
	fmt.Printf("s4:%s\n", s4)
}
