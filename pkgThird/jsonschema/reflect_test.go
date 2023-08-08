package jsonschema

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name      string  `json:"name" doc:"名字"`
	Age       int     `json:"age" form:"age" doc:"年龄"`
	MaxChild  Child   `json:"maxChild" form:"maxChild" doc:"最大的孩子"`
	ChildList []Child `json:"childList" form:"childList" doc:"所有的孩子"`
	School
}
type School struct {
	SchoolId uint `json:"schoolId" form:"schoolId" doc:"学校id"`
}

type Child struct {
	Name string `json:"name" form:"name" doc:"姓名"`
	Age  int    `json:"age" form:"age" doc:"年龄"`
}

func Test_fn(t *testing.T) {
	// res := DefaultRender{Data: User{}}
	res := User{}
	rt := reflect.TypeOf(&res)
	r := &Reflector{}
	s := r.ReflectFromType(rt)
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
}
