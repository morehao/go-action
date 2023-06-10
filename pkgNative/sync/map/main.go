package main

import (
	"fmt"
	"sync"
)

func main() {
	mapDemo()
}

func mapDemo() {
	var m sync.Map
	// 写入
	m.Store("a", 1)
	m.Store("b", 2)
	// 读取
	v, ok := m.Load("a")
	fmt.Println("load res:", v, ok)
	// 遍历
	m.Range(func(key, value interface{}) bool {
		fmt.Println("range res:", "key is ", key, "v is ", value)
		// 如果return false，则停止遍历
		return true
	})
	// 	读取or写入
	m.LoadOrStore("c", 3)
	c, cOk := m.Load("c")
	fmt.Println("LoadOrStore res:", c, cOk)
	// 	删除
	m.Delete("b")
	b, bOk := m.Load("b")
	fmt.Println("delete res:", b, bOk)
}
