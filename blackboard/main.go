package main

import "fmt"

func main() {
	fmt.Println(isValid("{}"))
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for _, v := range []byte(s) {
		if m[v] > 0 {
			if len(stack) == 0 || m[v] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return true
}
