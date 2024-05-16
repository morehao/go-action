package utils

import (
	"fmt"
	"testing"
)

func TestBitWiseOrToUint64BySlice(t *testing.T) {
	// 输出二进制
	fmt.Printf("%b\n", 1)
	fmt.Printf("%b\n", 2)
	fmt.Printf("%b\n", 3)
	fmt.Printf("%b\n", 4)
	// t.Log(1 << 1)
	// t.Log(2 << 1)
	// t.Log(3 << 1)
}

func TestBitwiseOrToSliceByUint64(t *testing.T) {
	n := uint64(7)
	t.Log(BitwiseOrToSliceByUint64(n))
}
