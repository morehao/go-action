package main

import "fmt"

func main() {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	fmt.Println("heapSort:", sortArray(arr))
}

/*
堆排序是指利用堆这种数据结构所设计的一种排序算法。堆是一个近似完全二叉树的结构，并同时满足堆的性质：即子结点的值总是小于（或者大于）它的父节点。
算法步骤:
1、将初始待排序列 (R1, R2, ……, Rn) 构建成大顶堆，此堆为初始的无序区；
2、将堆顶元素 R[1] 与最后一个元素 R[n] 交换，此时得到新的无序区 (R1, R2, ……, Rn-1) 和新的有序区 (Rn), 且满足 R[1, 2, ……, n-1]<=R[n]；
3、由于交换后新的堆顶 R[1] 可能违反堆的性质，因此需要对当前无序区 (R1, R2, ……, Rn-1) 调整为新堆，然后再次将 R [1] 与无序区最后一个元素交换，
得到新的无序区 (R1, R2, ……, Rn-2) 和新的有序区 (Rn-1, Rn)。
4、不断重复此过程直到有序区的元素个数为 n-1，则整个排序过程完成。
*/
func sortArray(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}
	// 构建堆顶
	for i := size/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, size)
	}
	// 调整堆结构+交换堆顶元素与末尾元素
	for i := size - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
	return nums
}

// 堆调整
func adjustHeap(nums []int, start int, end int) {
	temp := nums[start]
	for i := start*2 + 1; i < end; i = i*2 + 1 {
		if i+1 < end && nums[i] < nums[i+1] {
			i++
		}
		if nums[i] > temp {
			nums[start] = nums[i]
			start = i
		} else {
			break
		}
	}
	nums[start] = temp
}
