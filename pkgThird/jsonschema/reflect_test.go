package jsonschema

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name      string `json:"name" doc:"名字"`
	Age       int
	MaxChild  Child
	ChildList []Child
	School
}
type School struct {
	SchoolId uint `json:"schoolId" form:"schoolId" doc:"学校id"`
}

type Child struct {
	Name string `json:"name" form:"name" doc:"姓名"`
	Age  int    `json:"age" form:"age"`
}

func Test_fn(t *testing.T) {
	rt := reflect.TypeOf(&User{})
	r := &Reflector{}
	s := r.ReflectFromType(rt)
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}
