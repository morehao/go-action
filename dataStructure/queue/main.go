package main

import "fmt"

func main() {
	queue := Constructor()
	queue.Push(1)
	fmt.Println(queue.Pop())
	fmt.Println(queue.List)
	fmt.Println(queue.Peek())
}

type Queue struct {
	List []int
}

func Constructor() Queue {
	return Queue{
		List: []int{},
	}
}

func (q *Queue) Push(x int) {
	q.List = append(q.List, x)
}

func (q *Queue) Pop() int {
	if len(q.List) == 0 {
		return 0
	}
	item := q.List[0]
	q.List = q.List[1:]
	return item
}

func (q *Queue) Peek() int {
	if len(q.List) == 0 {
		return 0
	}
	return q.List[0]
}
