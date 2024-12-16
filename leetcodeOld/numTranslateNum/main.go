package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(12258))
}

/*
1、状态定义：dp[i]代表以x(i)结尾的数字的翻译方案数量
2、转移方程：如果x(i-1)和x(i)组成的数字可以被翻译，则d[i]= d[i-1]+d[i-2],否则d[i]=d[i-1].
能否被翻译是取决于组成的数字是否x>=10&&x<=25
3、初始状态：dp[0]=dp[1]=1
4、返回值：dp[i]
*/
func translateNum(num int) int {
	src := strconv.Itoa(num)
	prePreRes, preRes, res := 0, 0, 1
	for i := 0; i < len(src); i++ {
		prePreRes, preRes = preRes, res
		if i == 0 {
			continue
		}
		// 截取最后两位，判断是否能翻译
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			res += prePreRes
		}
	}
	return res
}
