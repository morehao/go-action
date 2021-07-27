package structStudy

import "fmt"

// NestedStruct 结构体嵌套
func NestedStruct() {
	type Person struct {
		name string
		sex  string
		age  int
	}
	type Student struct {
		Person
		id   int
		addr string
	}
	fmt.Printf("Student:%+v", Student{})
}

type person struct {
	name string
	city  string
	age  int
}

// NewPerson 构造函数
func NewPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

