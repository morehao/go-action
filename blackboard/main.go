package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i)
		fmt.Println(nums[i])
	}
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rowSize, columnSize := len(matrix), len(matrix[0])
	row, column := 0, columnSize-1
	for row < rowSize && column > 0 {
		num := matrix[row][column]
		if target == num {
			return true
		} else if target > num {
			column--
		} else {
			row++
		}
	}
	return false
}
