package main

import "fmt"

func main() {
	fn1()
	fn2()
}

func fn1() {
	var args int64 = 1
	modifiedNumber1(args) // args就是实际参数
	// 形参地址与实参地址不同，说明是变量的副本或变量的拷贝
	fmt.Printf("实际参数的地址 %p\n", &args)
	fmt.Printf("改动后的值是  %d\n", args)
}

func modifiedNumber1(args int64) { //这里定义的args就是形式参数
	fmt.Printf("形参地址 %p \n", &args)
	args = 10
}

func fn2() {
	var args int64 = 1
	addr := &args
	fmt.Printf("原始指针的内存地址是 %p\n", addr)
	fmt.Printf("指针变量addr存放的地址 %p\n", &addr)
	modifiedNumber2(addr) // args就是实际参数
	fmt.Printf("改动后的值是  %d\n", args)
}

func modifiedNumber2(addr *int64) { //这里定义的args就是形式参数
	// 形参addr的内存地址发生变化，所以不是引用传递；但是因为形参和实参都指向同一个内存地址，所以可以修改addr的值
	fmt.Printf("形参地址 %p \n", &addr)
	*addr = 10
}
