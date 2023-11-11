package main

import (
	"fmt"
	"github.com/morehao/go-action/utils"
	"reflect"
)

func main() {
	data := &Data{
		Items: []Item{
			{
				Price: 1.22245,
			},
		},
		Item: Item{
			Price: 1.22245,
		},
		ItemMap: map[string]Item{
			"1": {
				Price: 1.22245,
			},
		},
		NameMap: map[string][]string{
			"a": []string{},
		},
		PriceMap: map[string]float64{
			"1": 1.22245,
		},
	}
	responseFormat(reflect.ValueOf(data))
	fmt.Println(utils.ToJson(data))
}

type Data struct {
	Items []Item `json:"items"`
	Item
	ItemMap   map[string]Item     `json:"itemMap"`
	PriceList []float64           `json:"priceList"`
	NameList  []string            `json:"nameList"`
	NameMap   map[string][]string `json:"nameMap"`
	PriceMap  map[string]float64  `json:"priceMap"`
}

type Item struct {
	Price     float64   `json:"price" precision:"2"`
	PriceList []float64 `json:"priceList" precision:"2"`
	DescList  []string  `json:"descList"`
	// Children  []Item    `json:"children"`
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
			value := mapRange.Value()
			responseFormat(value)
		}

		// keys := val.MapKeys()
		// for _, v := range keys {
		// 	field := val.MapIndex(v)
		// 	responseFormat(field)
		// }
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			responseFormat(field)
		}
	case reflect.Ptr:
		if val.IsZero() {
			break
		}
		st := val.Elem()
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			responseFormat(field)
		}

	default:
		return
	}
}
