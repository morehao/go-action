package main

import (
	"fmt"
)

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.Push(-1)
	fmt.Println(minStack.Min())
}

/*
栈list中的最小元素始终对应栈minList的栈顶元素，即 min() 函数只需返回栈minList的栈顶元素即可。
*/

type MinStack struct {
	list    []int
	minList []int // 通过辅助栈降低复杂度
}

func Constructor() MinStack {
	return MinStack{
		list:    []int{},
		minList: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	if len(this.minList) == 0 {
		this.minList = append(this.minList, x)
	} else {
		minTop := this.minList[len(this.minList)-1]
		this.minList = append(this.minList, min(x, minTop))
	}
}

func (this *MinStack) Pop() {
	this.list = this.list[:len(this.list)-1]
	this.minList = this.minList[:len(this.minList)-1]
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) Min() int {
	return this.minList[len(this.minList)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
