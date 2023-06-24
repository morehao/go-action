package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(sortArr(arr))
}

/*
希尔排序也是一种插入排序，它是简单插入排序经过改进之后的一个更高效的版本，也称为递减增量排序算法。
希尔排序的基本思想是：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录 “基本有序” 时，再对全体记录进行依次直接插入排序。
算法步骤我们来看下希尔排序的基本步骤，在此我们选择增量 gap=length/2，缩小增量继续以 gap = gap/2 的方式，
这种增量选择我们可以用一个序列来表示，{n/2, (n/2)/2, ..., 1}，称为增量序列。
希尔排序的增量序列的选择与证明是个数学难题，我们选择的这个增量序列是比较常用的，也是希尔建议的增量，称为希尔增量，但其实这个增量序列不是最优的。
此处我们做示例使用希尔增量。
先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：
1、选择一个增量序列 {t1, t2, …, tk}，其中 (ti>tj, i<j, tk=1)；
2、按增量序列个数 k，对序列进行 k 趟排序；
3、每趟排序，根据对应的增量 t，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。仅增量因子为 1 时，
整个序列作为一个表来处理，表长度即为整个序列的长度。
*/
func sortArr(nums []int) []int {
	n := len(nums)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			currentItem := nums[i]
			preIndex := i - gap
			for preIndex >= 0 && nums[preIndex] > currentItem {
				nums[preIndex+gap] = nums[preIndex]
				preIndex -= gap
			}
			nums[preIndex+gap] = currentItem
		}
		gap = gap / 2
	}
	return nums
}
