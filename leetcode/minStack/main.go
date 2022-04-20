package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.Push(-1)
	fmt.Println(minStack.Min())
}

type MinStack struct {
	list    []int
	minList []int
}

// Constructor /** initialize your data structure here. */
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
		top := this.minList[len(this.minList)-1]
		this.minList = append(this.minList, min(x, top))
	}
}

func (this *MinStack) Pop() {
	if len(this.list) == 0 {
		return
	}
	this.list = this.list[:len(this.list)-1]
	this.minList = this.minList[:len(this.minList)-1]
}

func (this *MinStack) Top() int {
	if len(this.list) == 0 {
		return 0
	}
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
