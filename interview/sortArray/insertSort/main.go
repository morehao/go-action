package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(insertSortAsc(arr))
}

// 降序插入排序
func insertSortAsc(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
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
