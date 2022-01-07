package main

import "fmt"

func main() {
	res1 := add1(1, 2)
	fmt.Print(res1)
	res2 := add2(1, 2)
	fmt.Print(res2)
}

func add1(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()
	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func add2(x, y int) (z int) {
	defer func() {
		z = z + 4
	}()
	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}
