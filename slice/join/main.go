package main

import (
	"fmt"
	"strings"
)

func main() {
	strList := []string{"a", "b", "c"}
	joinStr := strings.Join(strList, ",")
	fmt.Println(joinStr)
}
