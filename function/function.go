package function

// 闭包
func Closure(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

// 递归-阶乘
func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}

// 递归-斐波那切数列

func Fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return Fibonaci(i-1) + Fibonaci(i-2)
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
