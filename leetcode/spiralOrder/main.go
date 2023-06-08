package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	total := len(matrix) * len(matrix[0])
	res := make([]int, total, total)
	count := 0
	// 左上坐标
	leftUp := []int{0, 0}
	// 右上坐标
	rightUp := []int{0, len(matrix[0]) - 1}
	// 左下坐标
	leftDown := []int{len(matrix) - 1, 0}
	// 右下坐标
	rightDown := []int{len(matrix) - 1, len(matrix[0]) - 1}
	for count < total {
		// 左上角到右上角,纵坐标变化
		for i := leftUp[1]; i <= rightUp[1] && count < total; i++ {
			res[count] = matrix[leftUp[0]][i]
			count++
		}
		// 左上和右上的横坐标+1
		leftUp[0]++
		rightUp[0]++
		// 	右上角到右下角，横坐标变化
		for i := rightUp[0]; i <= rightDown[0] && count < total; i++ {
			res[count] = matrix[i][rightDown[1]]
			count++
		}
		// 右上和右下的纵坐标-1
		rightUp[1]--
		rightDown[1]--
		// 	右下角到左下角，纵坐标变化
		for i := rightDown[1]; i >= leftDown[1] && count < total; i-- {
			res[count] = matrix[rightDown[0]][i]
			count++
		}
		// 右下和左下的横坐标-1
		rightDown[0]--
		leftDown[0]--
		// 从左下角到左上角，横坐标变化
		for i := leftDown[0]; i >= leftUp[0] && count < total; i-- {
			res[count] = matrix[i][leftDown[1]]
			count++
		}
		// 左下和左上的纵坐标 +1
		leftDown[1]++
		leftUp[1]++
	}
	return res
}
