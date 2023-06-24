package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(sortArray(arr))
}

func sortArray(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}
	max, min := math.MinInt, math.MaxInt
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	countList := make([]int, max-min+1)
	for _, v := range nums {
		countList[v-min]++
	}
	index := 0
	for i, v := range countList {
		for v > 0 {
			nums[index] = i + min
			v--
			index++
		}
	}
	return nums
}
