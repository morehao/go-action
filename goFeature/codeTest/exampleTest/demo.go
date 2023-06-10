package exampleTest

import "fmt"

func SayHello() {
	fmt.Println("Hello World")
}

func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

func PrintStudentName() {
	studentMap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}
	for _, v := range studentMap {
		fmt.Println(v)
	}
}
