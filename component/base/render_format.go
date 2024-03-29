package base

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ResponseFormat(data interface{}) {
	if data == nil {
		return
	}
	setFloat64Prec(data)
	responseFormat(reflect.ValueOf(data))
}

func responseFormat(val reflect.Value) {
	vType := val.Type()
	kd := val.Kind()
	switch kd {
	case reflect.Slice, reflect.Array:
		if val.IsNil() {
			if val.CanSet() {
				newSlice := reflect.MakeSlice(vType, 0, 0)
				val.Set(newSlice)
			}
		} else {
			for i := 0; i < val.Len(); i++ {
				field := val.Index(i)
				responseFormat(field)
			}
		}
	case reflect.Map:
		mapRange := val.MapRange()
		for mapRange.Next() {
			key := mapRange.Key()
			value := mapRange.Value()
			switch value.Kind() {
			case reflect.Ptr, reflect.Interface:
				if !value.IsNil() {
					responseFormat(value.Elem())
				}
			case reflect.Struct:
				newValue := reflect.New(value.Type()).Elem()
				newValue.Set(value)
				responseFormat(newValue.Addr())
				val.SetMapIndex(key, newValue)
			case reflect.Slice:
				if value.IsNil() {
					if value.CanSet() {
						newSlice := reflect.MakeSlice(vType, 0, 0)
						value.Set(newSlice)
					}
				} else {
					for i := 0; i < value.Len(); i++ {
						field := value.Index(i)
						responseFormat(field)
					}
				}
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			responseFormat(field)
		}
	case reflect.Ptr:
		if !val.IsNil() {
			st := val.Elem()
			for i := 0; i < st.NumField(); i++ {
				field := st.Field(i)
				responseFormat(field)
			}
		}
	default:
		return
	}
}

func isBasicDataType(kd reflect.Kind) bool {
	switch kd {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float64, reflect.Float32,
		reflect.String, reflect.Bool:
		return true
	}

	return false
}

func setFloat64Prec(v interface{}) {
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
				setFloat64Prec(field.Elem().Interface())
			}
		case reflect.Struct:
			setFloat64Prec(field.Addr().Interface())
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
					setFloat64Prec(elem.Addr().Interface())
				}
			}
		case reflect.Map:
			mapRange := field.MapRange()
			for mapRange.Next() {
				key := mapRange.Key()
				value := mapRange.Value()
				switch value.Kind() {
				case reflect.Ptr, reflect.Interface:
					if !value.IsNil() {
						setFloat64Prec(value.Elem().Interface())
					}
				case reflect.Struct:
					newValue := reflect.New(value.Type()).Elem()
					newValue.Set(value)
					setFloat64Prec(newValue.Addr().Interface())
					field.SetMapIndex(key, newValue)
				case reflect.Float64:
					precision, err := strconv.Atoi(precisionTag)
					if err != nil {
						fmt.Println("Invalid precision:", err)
						continue
					}
					rounded := round(value.Float(), precision)
					field.SetMapIndex(key, reflect.ValueOf(rounded))
				case reflect.Slice:
					for j := 0; j < value.Len(); j++ {
						elem := value.Index(j)
						if elem.Kind() == reflect.Float64 {
							precision, err := strconv.Atoi(precisionTag)
							if err != nil {
								fmt.Println("Invalid precision:", err)
								continue
							}
							rounded := round(elem.Float(), precision)
							elem.SetFloat(rounded)
						} else {
							setFloat64Prec(elem.Addr().Interface())
						}
					}
				}
			}

		}
	}
}

func round(x float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(x*pow) / pow
}
