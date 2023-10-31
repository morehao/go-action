package codeGenerator

import (
	"fmt"
	"reflect"
)

// GenStructTransferCode
// @Description: 生成结构体赋值的代码片段
// @param src 源结构体
// @param dest 目标结构体
// @return string 生成的代码片段
func GenStructAssignCode(src, dest interface{}) string {
	srcType := reflect.TypeOf(src)
	destType := reflect.TypeOf(dest)

	if srcType.Kind() != reflect.Struct || destType.Kind() != reflect.Struct {
		return "Both src and dest must be structs"
	}

	code := fmt.Sprintf("dest %s := %s{\n", destType.Name(), destType.Name())

	for i := 0; i < destType.NumField(); i++ {
		destField := destType.Field(i)
		srcField, found := srcType.FieldByName(destField.Name)
		if !found {
			continue // Skip fields not found in src
		}

		code += fmt.Sprintf("    %s: src.%s,\n", destField.Name, srcField.Name)
	}
	code += "}\n"

	return code
}
