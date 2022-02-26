package main

import "fmt"

func main() {
	res := mergeAlternately("abcde", "pqr")

	fmt.Println("-------------分割线----------")
	fmt.Println(res)
}

func mergeAlternately(str1, str2 string) string {
	res := make([]uint8, 0, len(str1)+len(str2))
	for i := 0; i < len(str1) || i < len(str2); i++ {
		if i < len(str1) {
			res = append(res, str1[i])
		}
		if i < len(str2) {
			res = append(res, str2[i])
		}
	}
	return string(res)
}
