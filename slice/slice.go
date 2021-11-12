package slice

import "fmt"

func LenAndCap() {
	// s1 := make([]int, 5)
	// fmt.Printf("The length of s1: %d\n", len(s1))
	// fmt.Printf("The capacity of s1: %d\n", cap(s1))
	// fmt.Printf("The value of s1: %d\n", s1)
	// s2 := make([]int, 5, 8)
	// fmt.Printf("The length of s2: %d\n", len(s2))
	// fmt.Printf("The capacity of s2: %d\n", cap(s2))
	// fmt.Printf("The value of s2: %d\n", s2)
	// s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// s4 := s3[3:6]
	// fmt.Printf("The length of s4: %d\n", len(s4))
	// fmt.Printf("The capacity of s4: %d\n", cap(s4))
	// fmt.Printf("The value of s4: %d\n", s4)
	s5 := []int{1, 2, 3}
	s5 = append(s5, []int{4}...)
	// 原始切片容量小于1024时，扩展后新切片的容量为原来2倍
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	s6 := make([]int, 1024)
	s6 = append(s6, 0)
	// 原始切片容量大于等于1024时，扩展后新切片的容量为原来1.25倍
	fmt.Printf("The capacity of s6: %d\n", cap(s6))
	s7 := []int{1, 2, 3}
	s7 = append(s7, []int{4, 5, 6, 7}...)
	// 扩展后新切片的容量比原容量2倍都大时，扩展后的新切片的容量会以新切片长度为基准
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
}

func ChangeSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2 = append(slice2, 4) // 超出slice2.cap限制，重新分配底层数组，原数组不受影响
	slice3 := slice1[0:2]
	slice3 = append(slice3, 0) // 未超出slice3.cap限制，重新分配底层数组，原数组也改变
	fmt.Println(slice1, slice2, slice3)
}

func AppendSlice() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	s = append(s, 100, 200)      // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}
