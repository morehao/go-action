package main

import "fmt"

func main() {
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(nums, 5))
}

// 线性查找,从右上角看，往左比右上角小，往右比右上角大
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	row, column := 0, columns-1
	for row < rows && column >= 0 {
		num := matrix[row][column]
		if num == target {
			return true
		} else if num > target {
			column--
		} else {
			row++
		}
	}
	return false
}
