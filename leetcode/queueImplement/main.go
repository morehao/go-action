package main

import "fmt"

func main() {
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
	fmt.Println(queue.DeleteHead())
}

type CQueue struct {
	inStack  []int
	outStack []int
}

/*
队列：先入先出，栈：先入后出。
将一个栈当作输入栈，用于压入appendTail 传入的数据；另一个栈当作输出栈，用于deleteHead 操作。
每次deleteHead 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序。
*/
func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		this.inToOut()
	}
	head := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return head
}

func (this *CQueue) inToOut() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}
