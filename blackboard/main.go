package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	// quickSort(arr, 0, len(arr)-1)
	// fmt.Println(arr)
	// fmt.Println(mergeSort(arr))
	// fmt.Println(selectSort(arr))
	// insertSort(arr)
	// fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

func quickSort(nums []int, start, end int) {
	if len(nums) < 2 || end <= start {
		return
	}
	i, j := start, end
	mid := nums[(start+end)/2]
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
	if i < end {
		quickSort(nums, i, end)
	}
	if j > start {
		quickSort(nums, start, j)
	}
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	i, j := 0, 0
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

func selectSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		tmp := nums[i]
		nums[i] = nums[min]
		nums[min] = tmp
	}
	return nums
}

func insertSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		j := i
		for j > 0 && nums[j-1] > current {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = current
	}
	return
}

func bubbleSort(nums []int) {
	if len(nums) == 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
