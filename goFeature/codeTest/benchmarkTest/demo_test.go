package benchmarkTest

import "testing"

func BenchmarkMakeSliceWithoutPreAlloc(b *testing.B) {
	// b.N是动态调整的
	for i := 0; i < b.N; i++ {
		MakeSliceWithoutPreAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithPreAlloc()
	}
}
