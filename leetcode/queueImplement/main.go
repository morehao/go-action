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

func (q *MyQueue) Push(item int) {
	q.inStack = append(q.inStack, item)
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.inToOut()
	}
	item := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return item
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.inToOut()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
func (q *MyQueue) inToOut() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}
