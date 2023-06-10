package main

import "fmt"

/*
s = s[low : high : max] 切片的三个参数的切片截取的意义为 low 为截取的起始下标（含），
high 为窃取的结束下标（不含 high），max 为切片保留的原切片的最大下标（不含 max）；
即新切片从老切片的 low 下标元素开始，len = high - low, cap = max - low；
high 和 max 一旦超出在老切片中越界，就会发生 runtime err，slice out of range。
另外如果省略第三个参数的时候，第三个参数默认和第二个参数相同，即 len = cap。
*/
func main() {
	a := []int{1, 2, 3, 4, 5}
	// 切割切片a生成切片b，使用相同的底层数组，只是将底层数组切割了一部分出来，但修改了b底层数组的首地址
	b := a[2:4]
	b[0] = 9
	fmt.Printf("a:%v\n", a)
	fmt.Printf("%p\n", &a[2])
	fmt.Printf("%p\n", &b[0])
	// 切片b，长度左闭右开，为2；宽度切割起始位置到末尾，为3
	fmt.Printf("b len:%d, cap:%d\n", len(b), cap(b))
}
