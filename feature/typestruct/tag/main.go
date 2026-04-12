package main

import (
	"fmt"
	"reflect"
)

type Child struct {
	Age string `json:"age"`
	Sex int    `json:"sex"`
}

func main() {
	t := Child{}
	ty := reflect.TypeOf(t)
	for i := 0; i < ty.NumField(); i++ {
		fmt.Printf("Field:%s, Tag:%s\n", ty.Field(i).Name, ty.Field(i).Tag)
	}
}
