package main

import "fmt"

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	str := []byte(s)
	res := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}
