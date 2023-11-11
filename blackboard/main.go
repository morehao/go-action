package main

import (
	"fmt"
	"github.com/morehao/go-action/utils"
	"math"
	"reflect"
	"strconv"
)

func main() {
	data := &Data{
		Num: 1.123456,
		NumList: []float64{
			1.123456,
		},
		List: []ListItem{
			{
				Num1: 1.123456,
			},
		},
		NumMap: map[string]float64{
			"num1": 1.123456,
			"num2": 1.123456,
		},
		NumListMap: map[string][]float64{
			"num1": []float64{
				1.123456,
				1.123456,
			},
		},
		ListItemMap: map[string]ListItem{
			"num1": {
				Num1: 1.123456,
			},
		},
		ListMap: map[string][]ListItem{
			"num1": []ListItem{
				{
					Num1: 1.123456,
				},
			},
		},
	}
	setFloat64Prec(data)
	fmt.Println(utils.ToJson(data))
}

type Data struct {
	Num         float64               `precision:"2"`
	NumList     []float64             `precision:"2"`
	List        []ListItem            `precision:"2"`
	NumMap      map[string]float64    `precision:"2"`
	NumListMap  map[string][]float64  `precision:"2"`
	ListItemMap map[string]ListItem   `precision:"2"`
	ListMap     map[string][]ListItem `precision:"2"`
}

type ListItem struct {
	Num1 float64 `precision:"2"`
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
