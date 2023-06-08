package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))

}

// 通过模拟压入、弹出
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		// 辅助栈栈顶元素和popped当前元素对比，相同则说明当前元素出栈顺序正确
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}
