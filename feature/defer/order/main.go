package main

import "fmt"

// defer函数在函数执行结束后执行，遵循先进后出原则
func main() {
	var res = 1
	defer fmt.Println(res)
	res = 2
	return
}
