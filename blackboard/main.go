package main

import "fmt"

func main() {
	fmt.Println(isValid(""))
}

func isValid(s string) bool {
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	var stack []byte
	for i := range s {
		if m[s[i]] > 0 {
			if stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
