package main

import "fmt"

func main() {

	// Creating slices
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice2 []int
	slice3 := make([]int, 5)

	// Before copying
	fmt.Println("------------before copy-------------")
	fmt.Printf("len=%-4d cap=%-4d slice1=%v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len=%-4d cap=%-4d slice2=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%-4d cap=%-4d slice3=%v\n", len(slice3), cap(slice3), slice3)

	// Copying the slices
	copy_1 := copy(slice2, slice1)
	fmt.Println()
	fmt.Printf("len=%-4d cap=%-4d slice1=%v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len=%-4d cap=%-4d slice2=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Println("Total number of elements copied:", copy_1)
}
