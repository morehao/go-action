package main

import "fmt"

func main() {
	// 要定义结构体的标准输出用String()，定义标准错误输出Error()，定义格式化输出Format()
	personInfo := person{
		Name: "张三",
		Age:  1,
	}
	fmt.Println(personInfo)
	fmt.Println(err{})
	fmt.Println(formatTest{})
}

type person struct {
	Name string
	Age  int64
}

func (s person) String() string {
	return fmt.Sprintf("我是%s, 今年%d岁", s.Name, s.Age)
}

type err struct {
}

func (e err) Error() string {
	return "我是err"
}

type formatTest struct {
}

func (f formatTest) Format(s fmt.State, c rune) {
	fmt.Printf("我是format")
}
