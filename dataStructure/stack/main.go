package main

import "fmt"

type MyStack struct {
	list []int
}

func Constructor() MyStack {
	return MyStack{
		list: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MyStack) Pop() int {
	if len(this.list) == 0 {
		return 0
	}
	item := this.list[len(this.list)-1]
	this.list = this.list[0 : len(this.list)-1]
	return item
}

func (this *MyStack) Top() int {
	if len(this.list) == 0 {
		return 0
	}
	return this.list[len(this.list)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.list) == 0
}

func main() {
	stock := Constructor()
	stock.Push(1)
	stock.Push(2)
	stock.Push(3)
	fmt.Println(stock.Pop())
	fmt.Println(stock.Top())
	fmt.Println(stock.Empty())
}
