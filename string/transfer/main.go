package main

import "fmt"

func main() {
	b := []byte{'H', 'e', 'l', 'l', 'o'}
	s := "Hello"
	fmt.Printf("byte to string: s:%s\n", byteToString(b))
	fmt.Printf("string to byte: s:%s\n", stringToByte(s))
	s1 := "你好"
	for i, v := range s1 {
		fmt.Printf("i:%d, v:%c\n", i, v)
	}
}

func byteToString(s []byte) string {
	return string(s)
}

func stringToByte(s string) []byte {
	return []byte(s)
}
