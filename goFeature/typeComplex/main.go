package main

import "fmt"

func main() {
	var x = complex(1, 2) // 1+2i
	var y = complex(3, 4) // 3+4i
	fmt.Println(x)
	fmt.Println(y)
	// TODO:复数运算待了解
	fmt.Println(x * y) // "(-5+10i)"
	// 获取实部
	fmt.Println(real(x * y)) // "-5"
	// 获取虚部
	fmt.Println(imag(x * y)) // "10"
}
