package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := nums[(start+end)/2]
	i, j := start, end
	for i <= j {
		for nums[i] < mid {
			i++
		}
		for nums[j] > mid {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if start < j {
		quickSort(nums, start, j)
	}
	if end > i {
		quickSort(nums, i, end)
	}
	return
}
