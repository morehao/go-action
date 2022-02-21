package main

import "fmt"

func main() {
	//	1、字面量初始化
	map1 := map[string]int{
		"a": 1,
	}
	fmt.Printf("map1:%v", map1)
	// 2、通过内置函数make初始化
	map2 := make(map[string]int)
	map2["a"] = 1
	fmt.Printf("map2:%v", map2)
	// 初始化map时可以指定容量
	map3 := make(map[string]int, 2)
	fmt.Printf("map3 len:%d", len(map3))
}
