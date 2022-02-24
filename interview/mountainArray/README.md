# 题目

符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：

`arr.length >= 3`

`存在 i（0 < i < arr.length - 1）`使得：

`arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]`


给定由整数组成的山峰数组 arr ，返回任何满足 `arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]` 的下标 i ，即山峰顶部。

