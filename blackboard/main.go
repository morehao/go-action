package main

import (
	"fmt"
)

// 计算最大公约数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	return a
}

func main() {
	fmt.Println(gcd(48, 18)) // 输出 6
	fmt.Println(gcd(17, 23)) // 输出 1
}
