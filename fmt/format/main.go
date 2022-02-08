package main

import (
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	fmtPrint()
	fmtPrintln()
	fmtPrintf()
}

func fmtPrint() {
	type Person struct {
		Name string
	}
	person := Person{
		Name: "zhangsan",
	}
	fmt.Print(person)
}

func fmtPrintln() {
	type Person struct {
		Name string
	}
	person := Person{
		Name: "zhangsan",
	}
	fmt.Println(person)
}

func fmtPrintf() {
	type Person struct {
		Name string
	}
	person := Person{
		Name: "zhangsan",
	}

	fmt.Printf("v-默认值格式输出(只输出值):%v\n", person)    // 默认值格式输出，只有值的输出,{zhangsan}
	fmt.Printf("+v-带有字段名的输出：%+v\n", person)       // 带有字段名的输出,{Name:zhangsan}
	fmt.Printf("#v-相应值的Go语法表示:%#v\n", person)     // 相应值的Go语法表示,main.Person={zhangsan}
	fmt.Printf("T-相应值的类型:%T\n", person)           // 相应值的类型的Go语法表示,main.Person
	fmt.Printf("b-二进制表示:%b\n", 5)                 // 二进制表示
	fmt.Printf("c-unicode码表示(输出中文):%c\n", 0x4E2d) // unicode码表示，输出中文
	fmt.Printf("d-十进制表示:%d\n", 18)                // 十进制表示
	fmt.Printf("s-输出字符串:%s\n", "Go语言")            // 输出字符串
	fmt.Printf("b-输出布尔值:%t\n", true)

	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:\n", err)
	}
	fmt.Printf("%+v\n", animals)

	var animals2 []Animal
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	json_iterator.Unmarshal(jsonBlob, &animals2)
	fmt.Printf("%+v\n", animals2)
}
