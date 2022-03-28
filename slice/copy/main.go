package main

import (
	"fmt"
	"unsafe"
)

func main() {
	shallowCopy()
	deepCopy()

}

func shallowCopy() {
	slice := []int{0, 1, 2, 3, 4}
	s1 := slice[:2]
	s2 := slice[:3]
	s3 := slice[:4]
	fmt.Println("shallowCopy-slice:", slice, "address:", unsafe.Pointer(&slice))
	fmt.Println("shallowCopy-s1:", s1, "address:", unsafe.Pointer(&s1))
	fmt.Println("shallowCopy-s2:", s2, "address:", unsafe.Pointer(&s2))
	fmt.Println("shallowCopy-s3:", s3, "address:", unsafe.Pointer(&s3))
}

func deepCopy() {
	slice := []int{0, 1, 2, 3, 4}
	s1 := make([]int, 0)
	s2 := make([]int, 4)
	s3 := make([]int, 5)
	copy1 := copy(s1, slice)
	copy2 := copy(s2, slice)
	// 拷贝前需要申请足够的空间
	copy3 := copy(s3, slice)
	fmt.Println("deepCopy-slice:", slice, "address:", unsafe.Pointer(&slice))
	fmt.Println("deepCopy-copy1:", copy1, "s1:", s1, "address:", unsafe.Pointer(&s1))
	fmt.Println("deepCopy-copy2:", copy2, "s2:", s2, "address:", unsafe.Pointer(&s2))
	fmt.Println("deepCopy-copy3:", copy3, "s3:", s3, "address:", unsafe.Pointer(&s3))
}
