package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aaabc"
	fmt.Println(compress(str))
}

func compress(str string) string {
	result, cnt := string(str[0]), 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			cnt++
			if i == len(str)-1 {
				result += fmt.Sprintf("%s", strconv.Itoa(cnt))
			}
		} else {
			result += fmt.Sprintf("%s", strconv.Itoa(cnt))
			result += fmt.Sprintf("%s", string(str[i]))
			cnt = 1
			if i == len(str)-1 {
				result += fmt.Sprintf("%s%s", string(str[i]), strconv.Itoa(1))
			}
		}
	}
	return result
}
