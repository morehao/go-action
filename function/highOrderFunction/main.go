package main

import "fmt"

func main() {
	op1 := func(x, y int) int {
		return x + y
	}
	add := Calculator1(op1)
	fmt.Print(add(1, 2))

	op2 := func(x, y int) int {
		return x + y
	}
	result, _ := Calculator2(1, 2, op2)
	fmt.Print(result)
}

// 高阶函数
type operate func(x, y int) int
type operatorFunc func(x, y int) (int, error)

func Calculator1(op operate) operatorFunc {
	if op == nil {
		return nil
	}
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, nil
		}
		return op(x, y), nil
	}
}

func Calculator2(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, nil
	}
	return op(x, y), nil
}
