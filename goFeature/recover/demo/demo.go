package demo

import "fmt"

// RecoverRepeat 两次recover 第二次无效
func RecoverRepeat() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("B")
		}
	}()

	panic("demo")
	fmt.Println("C")
}
