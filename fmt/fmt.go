/*
 * @Author: your name
 * @Date: 2021-05-24 12:35:08
 * @LastEditTime: 2021-05-24 12:38:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go-practice/fmt.go
 */

package fmt

import (
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main1() {
	//type Person struct {
	//	Name string
	//}
	//person := Person{
	//	Name: "zhangsan",
	//}
	//
	//fmt.Print(person)
	//fmt.Println(person, person)
	//fmt.Print(person)
	//// 默认值格式输出，只有值的输出,{zhangsan}
	//fmt.Printf("%v", person)
	//// 带有字段名的输出,{Name:zhangsan}
	//fmt.Printf("%+v", person)
	//// 相应值的Go语法表示,main.Person={zhangsan}
	//fmt.Printf("#v", person)
	//// 相应值的类型的Go语法表示,main.Person
	//fmt.Printf("%T", person)
	////	二进制表示
	//fmt.Printf("%b", 5)
	////	unicode码表示，输出中
	//fmt.Printf("%c", 0x4E2d)
	// 十进制表示
	//fmt.Printf("%d", 18)
	//// 输出字符串
	//fmt.Printf("%s", "Go语言")

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
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

	var animals2 []Animal
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	json_iterator.Unmarshal(jsonBlob, &animals2)
	fmt.Printf("%+v", animals2)

}
