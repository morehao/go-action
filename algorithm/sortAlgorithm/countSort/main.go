package main

import (
	"fmt"
	"math"
)

func main() {
	// var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	arr := []int{5, 3, 2, 2}
	fmt.Println(sortArray(arr))
}

/*
计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
计数排序 (Counting sort) 是一种稳定的排序算法。计数排序使用一个额外的数组 C，其中第 i 个元素是待排序数组 A 中值等于 i 的元素的个数。
然后根据数组 C 来将 A 中的元素排到正确的位置。它只能对整数进行排序。
算法步骤:
1、找出数组中的最大值 max、最小值 min；
2、创建一个新数组 C，其长度是 max-min+1，其元素默认值都为 0；
3、遍历原数组 A 中的元素 A[i]，以 A[i]-min 作为 C 数组的索引，以 A[i] 的值在 A 中元素出现次数作为 C[A[i]-min] 的值；
4、对 C 数组变形，新元素的值是该元素与前一个元素值的和，即当 i>1 时 C[i] = C[i] + C[i-1]；
5、创建结果数组 R，长度和原始数组一样。
6、从后向前遍历原始数组 A 中的元素 A[i]，使用 A[i] 减去最小值 min 作为索引，在计数数组 C 中找到对应的值 C[A[i]-min]，
C[A[i]-min]-1 就是 A[i] 在结果数组 R 中的位置，做完上述这些操作，将 count[A[i]-min] 减小 1。
*/
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
		// countList第 i 个元素是待排序数组nums中值等于 i 的元素的个数
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
