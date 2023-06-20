package main

import (
	"fmt"
)

func main() {
	fmt.Println(permutation("abc")) // [abc acb bac bca cba cab]

}

func permutation(s string) []string {
	var res []string
	sBytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(sBytes)-1 {
			res = append(res, string(sBytes))
		}
		dict := map[byte]bool{}
		for i := x; i < len(sBytes); i++ {
			// 1、过滤剪枝
			if !dict[sBytes[i]] {
				// 2、使用交换的方式维护当前路径下已使用的字母的未使用的字母，在sBytes[x]左边都是已经使用过的字母，就无需遍历了
				sBytes[x], sBytes[i] = sBytes[i], sBytes[x]
				dict[sBytes[x]] = true
				dfs(x + 1)
				// 3、撤销，恢复原状
				sBytes[x], sBytes[i] = sBytes[i], sBytes[x]
			}
		}
	}
	dfs(0)
	return res
}
