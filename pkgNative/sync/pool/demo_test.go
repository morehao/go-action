package main

import (
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		_ = json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		_ = json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
