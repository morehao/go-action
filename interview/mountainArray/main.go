package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 4, 2}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {
	for i := 1; ; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
}
