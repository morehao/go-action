package main

import (
	"fmt"
	"testing"
	"time"
)

// BenchmarkInterfaceAssignment 基准测试函数
func BenchmarkInterfaceAssignment(b *testing.B) {
	var a int = 3
	for i := 0; i < b.N; i++ {
		var i interface{} = a
		_ = i
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	expireAt := now.Add(time.Hour * 24 * 30)
	fmt.Println(expireAt)
}
