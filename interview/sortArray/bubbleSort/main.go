package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	bubbleSortAsc(arr)
	fmt.Println("bubbleSortAsc:", arr)
}

// 升序冒泡排序
func bubbleSortAsc(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			// 相邻两个元素比较大小，然后交换
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
