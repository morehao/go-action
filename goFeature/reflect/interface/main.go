package main

import (
	"fmt"
	"reflect"
)

// 空接口与反射
func main() {
	var x = 3.4
	reflectType(x)
	reflectValue(x)
	// 反射认为下面是指针类型，不是float类型
	reflectSetValue(&x)
}

// 反射获取interface类型信息
func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("a的类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println("kind()方法获取到a的类型是：", k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

// 反射获取interface值信息
func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("valueOf()的结果是：", v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}

// 反射修改值
func reflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}
