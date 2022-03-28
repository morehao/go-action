package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 8))
}

func search(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	// sort.SearchInts()函数在目标slice中搜索不到被搜索元素时,返回了被搜索的元素应该在目标slice中按升序排序该插入的位置
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}
