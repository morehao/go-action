package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int `json:"id"`
	Name string
	Age  int
}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

// 传入interface{},输出类型、字段和方法
func printInfo(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("field:%s, type:%v\n", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val:", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}

// 查看匿名结构体信息
func anonymousStructInfo(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t)
	// 匿名字段
	fmt.Printf("%#v\n", t.Field(0))
	// 非匿名字段
	fmt.Printf("%#v\n", t.Field(1))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(v).Field(0))
}

// 修改结构体值
func setValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}

func callFunc(u interface{}) {
	v := reflect.ValueOf(u)
	// 获取方法
	m := v.MethodByName("Hello")
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("6666")}
	// 没参数的情况下：var args2 []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args)
}

func getTag(s interface{}) {
	v := reflect.ValueOf(s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
}

func main() {
	// u := User{1, "zs", 20}
	// printInfo(u)
	// m := Boy{User{1, "zs", 20}, "bj"}
	// anonymousStructInfo(m)

	u1 := User{1, "5lmh.com", 20}
	// setValue(&u1)
	// fmt.Println(u1)
	// callFunc(u1) // TODO:调用报错，待修复
	getTag(&u1)
}
