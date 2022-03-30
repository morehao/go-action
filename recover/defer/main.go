package main

import "fmt"

func main() {
	correctDemo()
	incorrectDemo()
}

func correctDemo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("correctDemo recover err:", err)
		}
	}()
	panic("correctDemoPanic")
}

func recoverFn() {
	if err := recover(); err != nil {
		fmt.Println("incorrectDemo recover err:", err)
	}
}

// recover()函数必须被defer直接调用，间接调用无效
func incorrectDemo() {
	defer func() {
		recoverFn()
	}()
	panic("incorrectDemoPanic")
}
