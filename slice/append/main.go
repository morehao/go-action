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
	/*将切片作为参数传递给函数的时候，是值传递，传递给函数的其实是对切片底层结构体的拷贝；
	切片底层的value是指针，所以函数内部修改底切片，底层数组的值也会发生变化；
	切片底层结构体中长度和宽度并不是指针，函数内部改变切片的长度（未发生扩容)时其实底层数组的值已经改变了，
	但是因为是值传递，所以函数内部长度的改变并不会影响函数外切片结构体中长度的值，所以在函数外部看到的切片并未发生变化，*/
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
