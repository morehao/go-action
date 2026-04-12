package main

import (
	"fmt"
	"sort"
)

// 对map的key进行排序
func main() {
	var m = map[string]int{
		"hello":   0,
		"morning": 1,
		"my":      2,
		"girl":    3,
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
