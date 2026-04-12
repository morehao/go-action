package main

import (
	"fmt"
	"reflect"
)

func main() {
	firstLaw()
	secondLaw()
	thirdLaw()
}

// 第一定律：反射可以将interface类型变量转换成反射对象
func firstLaw() {
	var x = 3.4
	t := reflect.TypeOf(x) // t是反射对象reflect.Type
	fmt.Println("type:", t)
	v := reflect.ValueOf(x) // v是反射对象reflect.Value
	fmt.Println("value:", v)
}

// 第二定律：反射可以将反射对象还原成interface对象
func secondLaw() {
	var A interface{}
	A = 100
	v := reflect.ValueOf(A)
	B := v.Interface()
	if A == B {
		fmt.Println("They are same!")
	}
}

// 第三定律：反射对象可以修改，value值必须是可设置的
func thirdLaw() {
	var x = 3.4
	// reflect.ValueOf(x)是x的值，并非x本身，通过v修改值其实是无法影响x的，会报panic错误;
	// 通过反射可以修改interface的值，但是必须获得interface变量的地址；
	// 构建v时使用x的地址就可以修改x的值，同时通过Elem()可以获得执行value的指针。
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(3.14)
	fmt.Println("new x:", v.Elem().Interface())
}
