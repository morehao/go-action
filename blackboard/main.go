package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(sortArr(arr))
}

func sortArr(nums []int) []int {
	size := len(nums)
	if size < 2 {
		return nums
	}
	mid := size / 2
	return merge(sortArr(nums[:mid]), sortArr(nums[mid:]))
}

func merge(a, b []int) []int {
	i, j := 0, 0
	res := make([]int, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res[i+j] = a[i]
			i++
		} else {
			res[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		res[i+j] = a[i]
		i++
	}
	for j < len(b) {
		res[i+j] = b[j]
		j++
	}
	return res
}
