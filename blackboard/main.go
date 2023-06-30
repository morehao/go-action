package main

import "strings"

func main() {
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	return strings.Join(words, " ")
}
