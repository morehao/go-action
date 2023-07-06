package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1, nums2 := []int{1, 2, 3, 3}, []int{1, 3}
	fmt.Println(intersection(nums1, nums2))
}

// 双指针
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if len(res) == 0 || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return res
}
