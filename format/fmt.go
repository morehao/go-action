/*
 * @Author: your name
 * @Date: 2021-05-24 12:35:08
 * @LastEditTime: 2021-05-24 12:38:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go-practice/fmt.go
 */

package format

import (
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func FmtPrint() {
	type Person struct {
		Name string
	}
	person := Person{
		Name: "zhangsan",
	}
	fmt.Print(person)
}

func FmtPrintln() {
	type Person struct {
		Name string
	}
	person := Person{
		Name: "zhangsan",
	}
	fmt.Println(person)
}

func FmtPrintf() {
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
