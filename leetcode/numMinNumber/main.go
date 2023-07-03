package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{2, 10}
	minNumber(nums)
}

/*
此题求拼接起来的最小数字，本质上是一个排序问题。设数组nums 中任意两数字的字符串为x和y ，则规定 排序判断规则 为：
若拼接字符串x+y>y+x ，则x “大于”y ；反之，若x+y<y+x ，则x “小于”y .
x “小于”y 代表：排序完成后，数组中x 应在y 左边；“大于” 则反之。
*/
func minNumber(nums []int) string {
	var res strings.Builder
	sList := make([]string, 0, len(nums))
	for i := range nums {
		sList = append(sList, strconv.Itoa(nums[i]))
	}
	sort.Slice(sList, func(i, j int) bool {
		return sList[i]+sList[j] < sList[j]+sList[i]
	})
	for i := range sList {
		res.WriteString(sList[i])
	}
	return res.String()
}
