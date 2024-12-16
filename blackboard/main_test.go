package main

import "testing"

// BenchmarkInterfaceAssignment 基准测试函数
func BenchmarkInterfaceAssignment(b *testing.B) {
	var a int = 3
	for i := 0; i < b.N; i++ {
		var i interface{} = a
		_ = i
	}
}
