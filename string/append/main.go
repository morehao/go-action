package main

import "fmt"

func main() {
	// SlicePrint()
	a := []int{1, 2, 3}    //长度为3，容量为3
	b := make([]int, 1, 8) //长度为1，容量为8
	appendTest(a, b)
	fmt.Println(a)
	fmt.Println(b)
}

func SlicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2) // 输出[1 2] [2 3 4]
}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
func appendTest(a, b []int) {
	// 触发了扩容，连底层数组的地址都发生了改变，但是因为是值传递，所以在main中的a还是指向原来的底层数组，长度容量不变还是为3
	// test 函数中a触发了扩容，底层数组的地址和main种的a的底层数组不同了，长度为变为4，容量大于3了。
	a = append(a, 12)
	// 没有触发扩容，所以test种的b和main中的b指向了同一个底层数组，但是在test中b的长度变为了2，容量不变，因为没有扩容。
	// 因为是值传递，所以这个长度的变化并不能影响main中的b的切片结构体
	// 因为实质上b从main进入test还是值传递，只不过值拷贝的是b的切片结构体
	b = append(b, 12)
}
