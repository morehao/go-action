package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(selectionSort(arr))
}

/*
选择排序
1、首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置；
2、再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3、重复第 2 步，直到所有元素均排序完毕。
*/
func selectionSort(nums []int) []int {
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
