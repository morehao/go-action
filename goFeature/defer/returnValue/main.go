package main

import "fmt"

func main() {
	fmt.Println("unamed fn return：", unnamed())
	fmt.Println("named fn return：", named())
}

// 匿名返回值，返回值是return执行时声明的，defer无法访问，等价于return了一个i的值拷贝，defer修改i不影响返回值。
func unnamed() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer a:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer b :", i)
	}()
	return i // 0
}

// defer b : 1
// defer a: 2
// return: 0

// 具名返回值,具名返回值defer可以直接访问修改。
func named() (i int) {
	defer func() {
		i++
		fmt.Println("defer c:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer d :", i)
	}()
	return i // 2
}

// defer d : 1
// defer c: 2
// return: 2
