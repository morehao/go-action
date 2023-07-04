package main

import "fmt"

func main() {
	strs := []string{"aa", "bb", "cc"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < size; i++ {
		prefix = getPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func getPrefix(s1, s2 string) string {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	i := 0
	for i < minLen && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}
