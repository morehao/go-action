package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {

	queue := MyQueue{
		list: make([]int, 0),
	}
	result := make([]int, 0)
	// 构造第一个窗口
	for i := 0; i < k-1; i++ {
		queue.push(nums[i])
	}
	// 开始滑动窗口
	for i := k - 1; i < len(nums); i++ {
		// 右边新元素进入队列
		queue.push(nums[i])
		// 记录队列中最大元素
		result = append(result, queue.max())
		// 左边元素出队列
		queue.shift(nums[i-k+1])
	}
	return result
}

type MyQueue struct {
	list []int
}

func (q *MyQueue) push(v int) {
	// 删除list内所有小于v的值，确保list递减
	for len(q.list) > 0 && q.list[len(q.list)-1] < v {
		q.list = q.list[:len(q.list)-1]
	}
	q.list = append(q.list, v)
}

func (q *MyQueue) shift(v int) {
	if q.list[0] == v {
		q.list = q.list[1:]
	}
}

func (q *MyQueue) max() int {
	return q.list[0]
}
