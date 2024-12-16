package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 3
	var i interface{} = a

	fmt.Println(reflect.TypeOf(i)) // 输出：int
	fmt.Println(i)                 // 输出：3
}
