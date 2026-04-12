package main

import "fmt"

var TestMap map[string]int

func main() {
	MapCUDR()
	setKeyValue()
	getByKey("a")
}

func MapCUDR() {
	m := make(map[string]int)
	//	添加
	m["a"] = 1
	m["b"] = 2
	// 修改
	m["a"] = 11
	// 删除
	delete(m, "b")
	// 查询
	v, exist := m["a"]
	if exist {
		fmt.Println(v)
	}
}

// 未初始化的map为默认值为nil，向nil的map写入值会触发panic
func setKeyValue() {
	TestMap["a"] = 1
}

// 未初始化的map为默认值为nil，向nil的map读取值会返回对应值类型的0值
func getByKey(key string) int {
	a := TestMap["a"]
	return a
}
