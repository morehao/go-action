package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i)
		fmt.Println(nums[i])
	}
}

type MaxQueue struct {
	list    []int
	minList []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.list) == 0 {
		return -1
	}
	return this.minList[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.list) > 0 && this.list[len(this.list)-1] < value {
		this.list = this.list[:len(this.list)-1]
	}
	this.list = append(this.list, value)
	this.minList = append(this.minList, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.list) == 0 {
		return -1
	}
	if this.minList[0] == this.list[0] {
		this.minList = this.minList[1:]
	}
	v := this.list[0]
	this.list = this.list[1:]
	return v
}
