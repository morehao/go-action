package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	slideWindows := &HighQueue{
		List: []int{},
	}
	result := make([]int, 0)
	for i := 0; i < k-1; i++ {
		slideWindows.push(nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		slideWindows.push(nums[i])
		result = append(result, slideWindows.max())
		slideWindows.shift(nums[i-k+1])
	}
	return result
}

type HighQueue struct {
	List []int
}

func (q *HighQueue) push(v int) {
	for len(q.List) > 0 && q.List[len(q.List)-1] < v {
		q.List = q.List[:len(q.List)-1]
	}
	q.List = append(q.List, v)
}

func (q *HighQueue) shift(v int) {
	if v == q.List[0] {
		q.List = q.List[1:]
	}
}

func (q *HighQueue) max() int {
	if len(q.List) == 0 {
		return 0
	}
	return q.List[0]
}
