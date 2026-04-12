package main

func Slice() {
	// 栈空间不足时会发生逃逸
	s := make([]int, 10000, 10000)
	for i, _ := range s {
		s[i] = i
	}
}

func main() {
	Slice()
}
