package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

// 随机生成了 5 组数据，并且使用冒泡排序法排序
func main() {
	// 添加以下两行方法调用即可得到cpu性能数据
	//_ = pprof.StartCPUProfile(os.Stdout)
	//defer pprof.StopCPUProfile()

	// 一般来说，不建议将结果直接输出到标准输出，因为如果程序本身有输出，则会相互干扰，直接记录到一个文件中是最好的方式。
	f, _ := os.OpenFile("goFeature/pprof/cpu/cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer func() {
		_ = f.Close()
	}()
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
