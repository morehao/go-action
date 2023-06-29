package main

import "fmt"

func main() {
	fmt.Println(findContinuousSequence(9))
}

/*
因为是连续正整数，所以最小的和为3，即target为3，对应连续数组为[1,2];
要求序列长度至少大于2，所以枚举的上界为target/2；
*/
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	i, j, sum := 1, 2, 3
	for i <= target/2 {
		if target > sum {
			j++
			sum += j
		} else {
			if target == sum {
				tmp := make([]int, j-i+1)
				for k := i; k <= j; k++ {
					tmp[k-i] = k
				}
				res = append(res, tmp)
			}
			sum -= i
			i++
		}
	}
	return res
}
