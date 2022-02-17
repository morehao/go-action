package main

func Slice() []int {
	// 栈空间不足时会发生逃逸
	s := make([]int, 1000, 1000)
	for i, _ := range s {
		s[i] = i
	}
	return s
}

func main() {
	Slice()
}
