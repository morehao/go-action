package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("UnixNano:%d\n", time.Now().UnixNano())
	fmt.Printf("Unix:%v\n", time.Now().Unix())
	usec := uint64(time.Now().UnixNano())
	const (
		INT_MAX = 0x7FFFFFFF
		INT_MIN = 0x80000000
	)
	fmt.Println(strconv.FormatUint(usec&INT_MAX|INT_MIN, 10))
}
