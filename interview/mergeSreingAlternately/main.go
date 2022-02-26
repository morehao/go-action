package main

import "fmt"

func main() {
	res := mergeAlternately("abcde", "pqr")

	fmt.Println("-------------分割线----------")
	fmt.Println(res)
}

func mergeAlternately(str1, str2 string) string {
	maxStr, minStr := str1, str2
	if len(str1) < len(str2) {
		maxStr = str2
		minStr = str1
	}
	lenDiff := len(maxStr) - len(minStr)
	res := ""
	for i := 0; i < len(minStr); i++ {
		res = res + string(str1[i]) + string(str2[i])
	}
	if lenDiff > 0 {
		res += maxStr[len(maxStr)-lenDiff:]
	}
	return res
}
