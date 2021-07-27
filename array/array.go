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

func declare() {
	array1 := [3]int{}
	//我们可以通过如下方式给数组赋值
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	// array1[3] = 4 数组是大小固定的一种数据结构，所以该赋值异常
	//下面这种也是数组声明的一种方式，并且初始化数组的值为{1,2}
	array2 := [2]int {1,2}
	array3 := [2]int {4,5}
	array2 = array3
	// go语言中，数组是基本类型，数组之间是值传递，所以array3[0]重新赋值不会改变array2的最终值
	array3[0] = 6
	fmt.Println(array1, array2, array3)
}