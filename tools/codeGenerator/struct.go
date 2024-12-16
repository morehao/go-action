package codeGenerator

import (
	"fmt"
	"reflect"
)

// GenStructAssignCode 根据结构体字段生成从 src 到 dest 的赋值代码。
// src 和 dest 参数必须是结构体类型，并且该函数将通过字段名进行匹配。
// srcTag 和 destTag 参数可以用来指定生成代码中使用的变量名。
// 如果 srcTag 或 destTag 是空字符串，将分别使用默认变量名 "src" 或 "dest"。
// 函数返回包含生成的赋值代码的字符串。
// 如果 src 或 dest 任一不是结构体类型，函数将返回错误信息。
func GenStructAssignCode(src, dest interface{}, srcTag, destTag string) string {
	srcType := reflect.TypeOf(src)
	destType := reflect.TypeOf(dest)

	if srcType.Kind() != reflect.Struct || destType.Kind() != reflect.Struct {
		return "Both src and dest must be structs"
	}
	var srcName, destName = srcTag, destTag
	if srcTag == "" {
		srcName = "dest"
	}
	if destTag == "" {
		destName = "src"
	}
	code := fmt.Sprintf("%s := %s{\n", destName, destType.Name())

	for i := 0; i < destType.NumField(); i++ {
		destField := destType.Field(i)
		srcField, found := srcType.FieldByName(destField.Name)
		if !found {
			continue // Skip fields not found in src
		}

		code += fmt.Sprintf("    %s: %s.%s,\n", destField.Name, srcName, srcField.Name)
	}
	code += "}\n"

	return code
}
