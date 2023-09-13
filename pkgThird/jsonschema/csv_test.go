package jsonschema

import (
	"fmt"
	"testing"
)

type queryReq struct {
	Name string `json:"name" doc:"名字"`
	Age  int    `json:"age" form:"age" doc:"年龄"`
	subQueryReq
}

type subQueryReq struct {
	SchoolId uint64 `json:"schoolId" form:"schoolId" doc:"学校id"`
}

func Test_BuildCsv(t *testing.T) {
	csv := BuildCsv(&queryReq{})
	fmt.Println(csv)
}
