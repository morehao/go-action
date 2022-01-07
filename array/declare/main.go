package main

import "fmt"

// 数组的声明方式
func main() {
	array1 := [3]int{}
	// 我们可以通过如下方式给数组赋值
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	// array1[3] = 4 数组是大小固定的一种数据结构，所以该赋值异常
	// 下面这种也是数组声明的一种方式，并且初始化数组的值为{1,2}
	array2 := [2]int{1, 2}
	array3 := [2]int{4, 5}
	array2 = array3
	// go语言中，数组是基本类型，数组之间是值传递，所以array3[0]重新赋值不会改变array2的最终值
	array3[0] = 6
	fmt.Println(array1, array2, array3)
}
