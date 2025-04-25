package main

import (
	"encoding/json"
	"sync"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "wang", Age: 25})

func unmarshal() {
	stu := &Student{}
	_ = json.Unmarshal(buf, stu)
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func RunWithPool() {
	stu := studentPool.Get().(*Student)
	_ = json.Unmarshal(buf, stu)
	studentPool.Put(stu)
}
