package main

import (
	"fmt"
	"reflect"
)

func main() {
	print()
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age" form:"age"`
}

func print() {
	user := User{
		Name: "张三",
		Age:  18,
	}
	v := reflect.ValueOf(&user)
	v = v.Elem()
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
	fmt.Println(user)
}
