package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(strToInt("-91283472332"))
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
	intMax, intMin, boundary := math.MaxInt32, math.MinInt32, math.MaxInt32/10
	for _, v := range newStr[i:] {
		if !(v >= '0' && v <= '9') {
			break
		}
		// 情况一：执行拼接10×res≥2147483650越界
		// 情况二：拼接后是2147483648或2147483649越界，所以v>7也可以判断
		if res > boundary || (res == boundary && v > '7') {
			if intSign == 1 {
				return intMax
			} else {
				return intMin
			}
		}
		res = res*10 + int(v-'0')
	}
	return res * intSign
}
