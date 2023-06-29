package main

import "fmt"

// defer中函数参数在defer出现时就已确认，输出结果为：1
func main() {
	var res = 1
	defer fmt.Println(res)
	res = 2
	return
}
