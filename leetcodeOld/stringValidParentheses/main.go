package main

import "fmt"

func main() {
	s := "]}{["
	fmt.Println(isValid(s))
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
	stock := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if m[s[i]] > 0 {
			// 和栈顶元素对比
			if len(stock) == 0 || m[s[i]] != stock[len(stock)-1] {
				return false
			}
			stock = stock[:len(stock)-1]
		} else {
			// 先把左括号入栈
			stock = append(stock, s[i])
		}
	}
	return len(stock) == 0
}
