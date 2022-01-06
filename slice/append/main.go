package main

import "fmt"

func main() {
	fn1()
	fn2()
}

// 扩容时切片容量的变化
func fn1() {
	s1 := []int{1, 2, 3}
	s1 = append(s1, []int{4}...)
	// 原始切片容量小于1024时，扩展后新切片的容量为原来2倍
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	s2 := make([]int, 1024)
	s2 = append(s2, 0)
	// 原始切片容量大于等于1024时，扩展后新切片的容量为原来1.25倍
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	s3 := []int{1, 2, 3}
	s3 = append(s3, []int{4, 5, 6, 7}...)
	// 扩展后新切片的容量比原容量2倍都大时，扩展后的新切片的容量会以新切片长度为基准
	fmt.Printf("The capacity of s3: %d\n", cap(s3))
}

// 扩容时切片底层数组的变化
func fn2() {
	s1 := []int{1, 2}
	// s1和s2为同一个底层数组
	s2 := s1
	// s2的cap限制，重新分配底层数组，原底层数组不受影响
	s2 = append(s2, 3)
	fmt.Printf("1 s1:%v\n", s1)
	fmt.Printf("1 s2:%v\n", s2)
	// sliceHandle函数内对s1进行了扩容，产生了新的底层数组，所以+1后s1不受影响
	sliceHandle(s1)
	fmt.Printf("2 s1:%v\n", s1)
	fmt.Printf("2 s2 pointer address:%p\n", s2)
	sliceHandle(s2)
	fmt.Printf("2 s2:%v\n", s2)
	fmt.Printf("3 s2 pointer address:%p\n", s2)
}

func sliceHandle(s []int) {
	s = append(s, 0)
	for i, _ := range s {
		s[i] += 1
	}
	fmt.Printf("sliceHandle s:%v\n", s)
	fmt.Printf("sliceHandle s pointer address:%p\n", s)
}
