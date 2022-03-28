package main

func main() {

}

// 线性查找
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

// 暴力破解,TODO:待调试
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
