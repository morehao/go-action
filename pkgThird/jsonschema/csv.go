package jsonschema

import (
	"fmt"
	"reflect"
)

func BuildCsv(request interface{}) []string {
	// 判断request是否结构体
	t := reflect.TypeOf(request)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	// 如果不是结构体，则直接返回
	if t.Kind() != reflect.Struct {
		return nil
	}
	rv := reflect.ValueOf(request)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	var csvDocs []string
	for i := 0; i < t.NumField(); i++ {
		// 定义 参数名,类型,说明
		// 获取当前字段名
		fieldName := t.Field(i).Name
		// 定义固定占位字符串（是否必填,示例值,固定参数）
		// 获取当前字段的值
		// 获取当前字段的类型
		fieldKind := t.Field(i).Type.Kind()
		field := rv.Field(i)

		// 如果类型为结构体，则递归
		if fieldKind == reflect.Struct || (fieldKind == reflect.Ptr && field.Elem().Kind() == reflect.Struct) {
			var subCsvDocs []string
			if fieldKind == reflect.Ptr {
				subCsvDocs = BuildCsv(field.Interface())
			} else {
				subCsvDocs = BuildCsv(field.Addr().Interface())
			}
			// 如果返回值不为空，则拼接字符串
			if len(subCsvDocs) > 0 {
				csvDocs = append(csvDocs, subCsvDocs...)
			}
			continue
		}
		// 设置类型
		fieldType := toCsvFieldType(fieldKind)
		// 获取doc标签的值，设置paramDesc
		fieldDoc := rv.Type().Field(i).Tag.Get("doc")
		// 拼接字符串
		csvDocs = append(csvDocs, fmt.Sprintf("%s,%s,false,,,%s", fieldName, fieldType, fieldDoc))
	}
	return csvDocs
}

func toCsvFieldType(fieldKind reflect.Kind) (fieldType string) {
	switch fieldKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldType = "integer"
	case reflect.String:
		fieldType = "string"
	case reflect.Array, reflect.Slice:
		fieldType = "array"
	case reflect.Float32, reflect.Float64:
		fieldType = "number"
	default:
		fieldType = "string"
	}
	return fieldType
}
