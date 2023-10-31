package main

import (
	"reflect"
	"strconv"
)

func SetFloat64Prec1(data interface{}) error {
	v := reflect.ValueOf(data).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("precision")

		if tag != "" {
			prec, err := strconv.Atoi(tag)
			if err != nil {
				return err
			}

			switch field.Kind() {
			case reflect.Float64:
				newVal := round(field.Float(), prec)
				if err != nil {
					return err
				}
				field.SetFloat(newVal)
			case reflect.Slice, reflect.Array:
				for j := 0; j < field.Len(); j++ {
					item := field.Index(j)
					if item.Kind() == reflect.Float64 {
						newVal := round(item.Float(), prec)
						if err != nil {
							return err
						}
						item.SetFloat(newVal)
					} else if item.Kind() == reflect.Struct || item.Kind() == reflect.Slice || item.Kind() == reflect.Array || item.Kind() == reflect.Map {
						err := SetFloat64Prec1(item.Addr().Interface())
						if err != nil {
							return err
						}
					}
				}
			case reflect.Map:
				for _, key := range field.MapKeys() {
					value := field.MapIndex(key)
					if value.Kind() == reflect.Float64 {
						newVal := round(value.Float(), prec)
						if err != nil {
							return err
						}
						field.SetMapIndex(key, reflect.ValueOf(newVal))
					} else if value.Kind() == reflect.Struct || value.Kind() == reflect.Slice || value.Kind() == reflect.Array || value.Kind() == reflect.Map {
						err := SetFloat64Prec1(value.Interface())
						if err != nil {
							return err
						}
					}
				}
			case reflect.Struct:
				err := SetFloat64Prec1(field.Addr().Interface())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
