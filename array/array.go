package array

import "fmt"
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func test1() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

// 找出两数之和等于目标值的下标
func findTargetIndex(arr []int, target int) []int {
	indexMap := make(map[int]int, 0)
	for k, v := range arr {
		diff := target - v
		_, ok := indexMap[diff]
		if ok {
			return []int{k, indexMap[diff]}
		}
		indexMap[v] = k
	}
	return []int{}
}