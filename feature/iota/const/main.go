package main

import "fmt"

func main() {
	const (
		zero = iota
		one
		two
	)
	// const关键字出现时被重置为0
	const zero1 = iota
	fmt.Print(zero, one, two, zero1)
}
