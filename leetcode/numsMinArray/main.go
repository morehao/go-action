package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(nums))
}

// 二分法
/*
1.初始化： 声明i,j 双指针分别指向
nums 数组左右两端；
2.循环二分：设m=(i+j)/2为每次二分的中点（"/"代表向下取整除法，因此恒有i≤m<j),
可分为以下三种情况：
1.当nums[m]>nums[j]时：m一定在左排序数组中，即旋转点x一定在[m+1,j]闭区间内，
因此执行i=m+1;
2.当nums[m]<nums[j]时：m一定在右排序数组中，即旋转点x一定在[i,m]闭区间内，因此
执行j=m;
3.当nms[m]=nums[j]时：无法判断m在哪个排序数组中，即无法判断旋转点x在[i,m]还是
[m+1,j]区间中。解决方案：执行j=j-1缩小判断范围，分析见下文。
3.返回值：当i=j时跳出二分循环，并返回旋转点的值nums[i]即可。
*/
func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		// "/" 代表向下取整除法，因此恒有left ≤ m < j
		mid := (i + j) / 2
		if numbers[mid] < numbers[j] {
			j = mid
		} else if numbers[mid] > numbers[j] {
			i = mid + 1
		} else {
			// j=j−1 只需证明每次执行此操作后，旋转点 x 仍在 [i, j]区间内即可
			j--
		}
	}
	return numbers[i]
}

// 简单粗暴
func minArray2(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}
