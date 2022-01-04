package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		// recover确保程序不宕掉
		if p := recover(); p != nil {
			fmt.Printf("panic:%s", p)
		}
		fmt.Println("defer function end")
	}()
	// 引发panic
	panic(errors.New("something wrong"))
}
