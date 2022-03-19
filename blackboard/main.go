package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", fib(5))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre, current, i := 0, 1, 2
	for i <= n {
		next := pre + current
		pre = current
		current = next
		i++
	}
	return current
}

func lengthOfLongestSubstring(s string) int {
	ans, rk := 0, -1
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < len(s) && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk+1-i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func compressString(s string) string {
	res := make([]byte, 0, len(s))
	i, j := 0, 0
	sLen := len(s)
	for i < sLen {
		for j < sLen && s[j] == s[i] {
			j++
		}
		res = append(res, s[i])
		res = append(res, []byte(strconv.Itoa(j-i))...)
		if len(res) >= sLen {
			return s
		}
		i = j
	}
	return string(res)
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if target < num {
			high = mid - 1
		} else if target == num {
			return mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
