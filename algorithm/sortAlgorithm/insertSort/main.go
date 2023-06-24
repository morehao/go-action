package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(insertSortAsc(arr))
}

/*
插入排序
1、从第一个元素开始，该元素可以认为已经被排序；
2、取出下一个元素，在已经排序的元素序列中从后向前扫描；
3、如果该元素（已排序）大于新元素，将该元素移到下一位置；
4、重复步骤 3，直到找到已排序的元素小于或者等于新元素的位置；
5、将新元素插入到该位置后；
6、重复步骤 2~5。
*/
func insertSortAsc(nums []int) []int {
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
