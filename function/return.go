package function

func Add1(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()
	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func Add2(x, y int) (z int) {
	defer func() {
		z = z + 4
	}()
	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

