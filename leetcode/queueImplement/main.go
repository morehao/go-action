package main

import "fmt"

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[0 : len(this.inStack)-1]
		}
	}
	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[0 : len(this.inStack)-1]
		}
		return this.outStack[len(this.outStack)-1]
	}
	x := this.outStack[len(this.outStack)-1]
	return x
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0
}
