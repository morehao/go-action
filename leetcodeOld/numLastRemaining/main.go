package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(5, 3))
}

/*
我们将上述问题建模为函数f(n,m),该函数的返回值为最终留下的元素的序号。
首先，长度为n的序列会先删除第m%n个元素，然后剩下一个长度为n-1的序列。
那么，我们可以递归地求解f(n-1,m),就可以知道对于剩下的n-1个元素，最终会留下第几个元素，我们设答案为x=f(n-1,m)。
由于我们删除了第m号n个元素，将序列的长度变为n-1。当我们知道了f(n-1,m)对应的答案x之后，我们也就可以知道，
长度为n的序列最后一个删除的元素，应当是从m%n开始数的第x个元素。因此有f(n,m)=(m%n+x)%n=(m+x)%n。
*/
func lastRemaining(n int, m int) int {
	return f(n, m)
}

func f(n, m int) int {
	if n == 1 {
		return 0
	}
	x := f(n-1, m)
	return (m + x) % n
}
