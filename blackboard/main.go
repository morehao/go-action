package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{10, 2}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}
