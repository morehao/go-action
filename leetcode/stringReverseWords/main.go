package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue!"
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	words := strings.Fields(s)
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	return strings.Join(words, " ")
}
