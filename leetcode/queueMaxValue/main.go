package main

func main() {

}

// 维护一个单调的双端队列DQueue
type MaxQueue struct {
	Data  []int
	Deque []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.Deque) > 0 && this.Deque[len(this.Deque)-1] < value {
		this.Deque = this.Deque[:len(this.Deque)-1]
	}
	this.Deque = append(this.Deque, value)
	this.Data = append(this.Data, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Data) == 0 {
		return -1
	}
	if this.Data[0] == this.Deque[0] {
		this.Deque = this.Deque[1:]
	}
	value := this.Data[0]
	this.Data = this.Data[1:]
	return value
}
