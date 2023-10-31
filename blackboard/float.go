package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func SetFloat64Prec(v interface{}) {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)

		precisionTag := typeField.Tag.Get("precision")
		if precisionTag != "" && field.Kind() == reflect.Float64 {
			if !field.CanSet() {
				fmt.Println("Cannot set field:", typeField.Name)
				continue
			}

			precision, err := strconv.Atoi(precisionTag)
			if err != nil {
				fmt.Println("Invalid precision:", err)
				continue
			}

			rounded := round(field.Float(), precision)
			field.SetFloat(rounded)
		}

		// Handle nested fields
		switch field.Kind() {
		case reflect.Ptr, reflect.Interface:
			if !field.IsNil() {
				SetFloat64Prec(field.Elem().Interface())
			}
		case reflect.Struct:
			SetFloat64Prec(field.Addr().Interface())
		case reflect.Slice:
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j)
				if elem.Kind() == reflect.Float64 {
					precision, err := strconv.Atoi(precisionTag)
					if err != nil {
						fmt.Println("Invalid precision:", err)
						continue
					}
					rounded := round(elem.Float(), precision)
					elem.SetFloat(rounded)
				} else {
					SetFloat64Prec(elem.Addr().Interface())
				}
			}
		case reflect.Map:
			for _, key := range field.MapKeys() {
				value := field.MapIndex(key)
				if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
					SetFloat64Prec(value.Elem().Interface())
				}
			}
		}
	}
}

func round(x float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(x*pow) / pow
}
