package main

import "fmt"

func main() {
	s := []string{"monkey"}
	fmt.Printf("findMonkey: %v", findMonkey(s))
	rangeDemo()
}

func findMonkey(s []string) bool {
	for _, v := range s {
		// 替换为v[i] == "monkey",性能会提升，因为for-range遍历过程中每次迭代均会执行一次赋值操作，赋值操作会涉及内存拷贝，性能上不如使用切片下标
		if v == "monkey" {
			return true
		}
	}
	return false
}

func rangeDemo() {
	s := []int{1, 2, 3}
	for i := range s {
		s = append(s, i)
	}
	fmt.Printf("s:%v", s)
}
