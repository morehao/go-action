package main

import "fmt"

func main() {
	fmt.Println(mySqrt(10))
	fmt.Println(cubeRoot(27, 0.1))
}

func mySqrt(x int) int {
	l, r := 0, x
	var ans int
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 求近似立方根
func cubeRoot(n, e float64) float64 {
	left, right := float64(1), n
	var mid float64
	for right-left > e {
		mid = (left + right) / 2
		cube := mid * mid * mid
		if cube > n {
			right = mid
		} else if cube < n {
			left = mid
		} else {
			return mid
		}
	}
	return mid
}
