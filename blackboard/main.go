package main

import (
	"fmt"
)

func main() {
	res := findTargetIndex([]int{1, 3, 5, 8, 7}, 3)
	fmt.Printf("result:%v", res)
}

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() *MyQueue {
	return &MyQueue{
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

func findTargetIndex(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := m[diff]; ok {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return nil
}

func search(nums []int, item int) int {
	i, j := 0, len(nums)
	for i <= j {
		mid := (j-i)/2 + i
		if item == nums[mid] {
			return mid
		} else if item < nums[mid] {
			j = mid
		} else {
			i = mid
		}
	}
	return -1
}

func bottomIndexInMountainArray(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return i
		}
	}
	return -1
}

func plusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = (nums[i] + 1) % 10
		if nums[i] != 0 {
			return nums
		}
	}
	newNums := make([]int, len(nums)+1)
	newNums[0] = 1
	return newNums

}

func bubbleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func selectSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		tmp := nums[i]
		nums[i] = nums[min]
		nums[min] = tmp

	}
}

func insertSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		j := i
		for j > 0 && nums[j-1] > current {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = current
	}
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	mid := nums[(start+end)/2]
	for i <= j {
		for nums[i] < mid {
			i++
		}
		for nums[j] > mid {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if i < end {
		quickSort(nums, i, end)
	}
	if j > start {
		quickSort(nums, start, j)
	}

}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[i+j] = left[i]
			i++
		} else {
			res[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		res[i+j] = left[i]
		i++
	}
	for j < len(right) {
		res[i+j] = right[j]
		j++
	}
	return res
}
