package benchmarkTest

func MakeSliceWithoutPreAlloc() []int {
	var s []int
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
	return s
}

func MakeSliceWithPreAlloc() []int {
	var s []int
	s = make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
	return s
}
