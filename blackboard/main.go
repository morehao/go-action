package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "43"
	fmt.Println(strToInt(s))
}

func strToInt(str string) int {
	newStr := strings.TrimSpace(str)
	if len(newStr) == 0 {
		return 0
	}
	res, i, intSign := 0, 1, 1
	if newStr[0] == '-' {
		intSign = -1
	} else if newStr[0] != '+' {
		i = 0
	}
	max, min, boundary := math.MinInt32, math.MinInt32, math.MinInt32/10
	for _, v := range newStr[i:] {
		if !(v >= '0' && v <= '9') {
			break
		}
		if res > boundary || (res == boundary && v > '7') {
			if intSign == 1 {
				return max
			} else {
				return min
			}
		}
		res = res*10 + int(v-'0')
	}
	return intSign * res
}
