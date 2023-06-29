package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(6))
}

/*
我们可以直接循环检查给定整数的二进制位的每一位是否为1。
具体代码中，当检查第i位时，我们可以让n与2^i进行与运算，当且仅当n的第i位为1时，运算结果不为0。
*/
func hammingWeight(num uint32) int {
	total := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			total++
		}
	}
	return total
}
