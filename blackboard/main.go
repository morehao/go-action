package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	ReverseList[int](s)
	fmt.Print(s)
}

func ReverseList[T any](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
