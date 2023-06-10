package main

import "fmt"

// defer函数在函数执行结束后执行，遵循先进后出原则
func main() {
	// defer func() {
	// 	fmt.Println(1)
	// }()
	// defer func() {
	// 	fmt.Println(2)
	// }()
	// defer func() {
	// 	fmt.Println(3)
	// }()
	// fmt.Printf("main function \n")
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}
