package main

import "fmt"

// 类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型。
func main() {
	assertSuccess()
	assertFailed()
	// assertPanic()
	getType(10)
}

func assertSuccess() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println("assertSuccess ", "value:", value, "ok:", ok)
}

func assertFailed() {
	var x interface{}
	x = 10
	value, ok := x.(struct{})
	fmt.Println("assertFailed ", "value:", value, "ok:", ok)
}

func assertPanic() {
	var x interface{}
	x = "Hello"
	// x = nil
	// 需要注意如果不接收第二个参数也就是上面代码中的 ok，断言失败时会直接造成一个 panic。如果 x 为 nil 同样也会 panic。
	value := x.(int)
	fmt.Println("assertPanic ", "value:", value)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
