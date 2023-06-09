package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var (
		total     = len(matrix) * len(matrix[0])
		count     = 0
		leftUp    = []int{0, 0}
		rightUp   = []int{0, len(matrix[0]) - 1}
		leftDown  = []int{len(matrix) - 1, 0}
		rightDown = []int{len(matrix) - 1, len(matrix[0]) - 1}
		res       = make([]int, total, total)
	)
	for count < total {
		// 	从左上角到右上角，纵坐标移动,移动完成之后，横坐标下移，即+1
		for i := leftUp[1]; i <= rightUp[1] && count < total; i++ {
			res[count] = matrix[leftUp[0]][i]
			count++
		}
		leftUp[0]++
		rightUp[0]++
		// 	从右上角到右下角，横坐标移动，移动完成之后，纵坐标左移，即-1
		for i := rightUp[0]; i <= rightDown[0] && count < total; i++ {
			res[count] = matrix[i][rightDown[1]]
			count++
		}
		rightUp[1]--
		rightDown[1]--
		// 	从右下角到左下角，纵坐标移动，移动完成之后横坐标上移，即-1
		for i := rightDown[1]; i >= leftDown[1] && count < total; i-- {
			res[count] = matrix[leftDown[0]][i]
			count++
		}
		rightDown[0]--
		leftDown[0]--
		// 	从左下角到左上角，横坐标移动，移动完成后纵坐标右移，即+1
		for i := leftDown[0]; i >= leftUp[0] && count < total; i-- {
			res[count] = matrix[i][leftUp[1]]
			count++
		}
		leftDown[1]++
		leftUp[1]++
	}
	return res
}
