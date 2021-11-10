package function

import (
	"fmt"
	"testing"
)

func Test_Add2(t *testing.T) {
	res := Add2(1, 2)
	fmt.Print(res)
}

func Test_Closure(t *testing.T) {
	f1, f2 := Closure(10)
	// base初始值为10
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}

func Test_Factorial(t *testing.T) {
	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
}

func Test_Fibonaci(t *testing.T) {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", Fibonaci(i))
	}
}

func Test_Defer1(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		// for循环中调用defer可能存在资源泄露
		defer t.Close()
	}
}

// 正常输出
func Test_Defer2(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer Close(t)
	}
	// 或
	// for _, t := range ts {
	// 	t2 := t
	// 	defer t2.Close()
	// }
}

func Test_Calculator1(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}
	add := Calculator1(op)
	fmt.Print(add(1, 2))
}

func Test_Calculator2(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}
	result, _ := Calculator2(1, 2, op)
	fmt.Print(result)
}
