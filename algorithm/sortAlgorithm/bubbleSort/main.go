package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	bubbleSortAsc(arr)
	fmt.Println("bubbleSortAsc:", arr)
}

/*
冒泡排序
1、比较相邻的元素。如果第一个比第二个大，就交换它们两个；
2、对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
3、针对所有的元素重复以上的步骤，除了最后一个；
4、重复步骤 1~3，直到排序完成。
*/
func bubbleSortAsc(nums []int) []int {
	for range nums {
		for j := 0; j < len(nums)-1; j++ {
			// 相邻两个元素比较大小，然后交换
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
