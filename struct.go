package main

import "fmt"

type student struct{
	name string
	age  int
}
func main(){
	arr1 := []int{1, 2, 3}
	arr2 := make([]*int, len(arr1))

	for i, v := range arr1 {
		arr2[i] = &v
	}

	for _, v := range arr2 {
		fmt.Println(*v)
	}
}