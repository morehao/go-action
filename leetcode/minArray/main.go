package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(nums))
}

// 二分法
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		// "/" 代表向下取整除法，因此恒有left ≤ m < right
		mid := (left + right) / 2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			// right=right−1 只需证明每次执行此操作后，旋转点 x 仍在 [left, right]区间内即可
			right--
		}
	}
	return numbers[left]
}

// 简单粗暴
func minArray2(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}
