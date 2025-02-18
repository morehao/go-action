package main

import (
	"fmt"
)

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(mergeSort(arr))
}
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func inertSort(nums []int) []int {
	for i := range nums {
		currentItem := nums[i]
		j := i
		for j > 0 && nums[j-1] > currentItem {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = currentItem
	}
	return nums
}

func selectSort(nums []int) []int {
	for i := range nums {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func quickSort(nums []int) []int {
	return quickPartition(nums, 0, len(nums)-1)
}

func quickPartition(nums []int, start, end int) []int {
	i, j := start, end
	midValue := nums[(start+end)/2]
	for i <= j {
		for nums[i] < midValue {
			i++
		}
		for nums[j] > midValue {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if i < end {
		quickPartition(nums, i, end)
	}
	if j > start {
		quickPartition(nums, start, j)
	}
	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	midIndex := len(nums) / 2
	a := mergeSort(nums[:midIndex])
	b := mergeSort(nums[midIndex:])
	return merge(a, b)
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
