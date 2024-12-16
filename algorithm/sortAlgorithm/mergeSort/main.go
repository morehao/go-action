package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println(sortArray(arr))
}

/*
归并排序算法是一个递归过程，边界条件为当输入序列仅有一个元素时，直接返回，具体过程如下：
1、如果输入内只有一个元素，则直接返回，否则将长度为 n 的输入序列分成两个长度为 n/2 的子序列；
2、分别对这两个子序列进行归并排序，使子序列变为有序状态；
3、设定两个指针，分别指向两个已经排序子序列的起始位置；
4、比较两个指针所指向的元素，选择相对小的元素放入到合并空间（用于存放排序结果），并移动指针到下一位置；
5、重复步骤 3 ~4 直到某一指针达到序列尾；
6、将另一序列剩下的所有元素直接复制到合并序列尾。
*/
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	a := sortArray(nums[:mid])
	b := sortArray(nums[mid:])
	return merge(a, b)
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	var (
		i = 0
		j = 0
	)
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
