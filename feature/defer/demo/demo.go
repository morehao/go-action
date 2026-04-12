package demo

import "fmt"

func DeferOrder() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func DeferFnParams() {
	var i = 1

	defer func(i int) {
		fmt.Println(i)
	}(i)
	// 等同于：defer fmt.Println(i)
	i = 2
	return
}

func DeferReturnValue1() {
	var i = 0
	defer func() {
		fmt.Println(i)
	}()
	i++
	// 	绑定了变量i
}

func DeferReturnValue2() {
	var arr = [3]int{1, 2, 3}
	defer func(array *[3]int) {
		for i := range array {
			fmt.Print(array[i])
		}
	}(&arr) // 绑定了底层数组地址
	arr[0] = 10
}

func DeferReturnValue3() (res int) {
	i := 1

	defer func() {
		res++
	}()

	return i
}

// DeferNest defer嵌套
func DeferNest() {
	defer func() {
		defer func() {
			fmt.Print("B")
		}()
		fmt.Print("A")
	}()
}
