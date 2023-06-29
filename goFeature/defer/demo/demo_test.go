package demo

import (
	"fmt"
)

// defer执行顺序后进先出
func ExampleDeferOrder() {
	DeferOrder()
	// Output:
	// 43210
}

func ExampleDeferFnParams() {
	DeferFnParams()
	// Output:
	// 1
}

func ExampleDeferReturnValue1() {
	DeferReturnValue1()
	// Output:
	// 1
}

func ExampleDeferReturnValue2() {
	DeferReturnValue2()
	// Output:
	// 1023
}

func ExampleDeferReturnValue3() {
	res := DeferReturnValue3()
	fmt.Print(res)
	// Output:
	// 2
}

func ExampleDeferNest() {
	DeferNest()
	// Output:
	// AB
}
