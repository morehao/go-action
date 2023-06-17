package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

/*
思路：深度优先搜索（DFS）+ 剪枝
深度优先搜索： 可以理解为暴力法遍历矩阵中所有字符串可能性。DFS 通过递归，先朝一个方向搜到底，再回溯至上个节点，沿另一个方向搜索，以此类推。
剪枝： 在搜索中，遇到 这条路不可能和目标字符串匹配成功 的情况（例如：此矩阵元素和目标字符不同、此元素已被访问），则应立即返回，称之为 可行性剪枝 。
*/
func exist(board [][]byte, word string) bool {
	var dfsCheck func(i, j, k int) bool
	dfsCheck = func(i, j, k int) bool {
		// 越界，返回true
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}
		// 当前字符不匹配，返回false
		if board[i][j] != word[k] {
			return false
		}
		// 已经访问到字符串的末尾，且对应字符依然匹配，返回true
		if k == len(word)-1 {
			return true
		}
		// 标记当前矩阵元素： 将 board[i][j] 修改为 空字符 '' ，代表此元素已访问过，防止之后搜索时重复访问。
		board[i][j] = ' '
		// 朝当前元素上下左右四个方向开始递归查找
		res := dfsCheck(i+1, j, k+1) || dfsCheck(i-1, j, k+1) || dfsCheck(i, j+1, k+1) || dfsCheck(i, j-1, k+1)
		// 还原当前矩阵元素
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfsCheck(i, j, 0) {
				return true
			}
		}
	}
	return false
}
