package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println("quickAscendingSort:", sortArray(arr))
}

/*
快速排序使用分治法策略来把一个序列分为较小和较大的 2 个子序列，然后递回地排序两个子序列。具体算法描述如下：
1、从序列中随机挑出一个元素，做为 “基准”(pivot)；
2、重新排列序列，将所有比基准值小的元素摆放在基准前面，所有比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个操作结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3、递归地把小于基准值元素的子序列和大于基准值元素的子序列进行快速排序。
*/
func sortArray(nums []int) []int {
	partition(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, start, end int) {
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
		partition(nums, i, end)
	}
	if j > start {
		partition(nums, start, j)
	}
	return
}
