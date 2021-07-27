package slice

import "fmt"

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
	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}