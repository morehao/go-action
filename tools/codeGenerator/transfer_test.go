package codeGenerator

import (
	"fmt"
	"testing"
)

func TestGenStructAssignCode(t *testing.T) {
	type SourceStruct struct {
		Field1 int
		Field2 string
	}

	type DestinationStruct struct {
		Field1 int
		Field2 string
		Field3 bool
	}
	var src SourceStruct
	var dest DestinationStruct
	code := GenStructAssignCode(src, dest)
	fmt.Println(code)
}
