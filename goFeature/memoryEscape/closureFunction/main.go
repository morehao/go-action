package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("fibonacci: %d\n", f())
	}
}

// 递归-斐波那切数列
// fibonacci()函数中原本属于局部变量的a和b由于闭包的引用，不得不将两个变量放到堆中，导致发生逃逸
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
