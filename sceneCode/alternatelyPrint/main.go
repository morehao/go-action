package main

import (
	"fmt"
)

/*
1、问题描述
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
2、解题思路
问题很简单，使用 channel 来控制打印的进度。使用两个 channel ，来分别控制数字和字母的打印序列， 数字打印完成后通过 channel 通知字母打印, 字母打印完成后通知数字打印，然后周而复始的工作。
*/
func main() {
	numberCh, letterCh, doneCh := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go func() {
		i := 1
		for range numberCh {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letterCh <- struct{}{}
		}
	}()
	go func() {
		i := 'A'
		for range letterCh {
			if i >= 'Z' {
				doneCh <- struct{}{}
			} else {
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				numberCh <- struct{}{}
			}
		}
	}()
	numberCh <- struct{}{}
	// for range doneCh 等同于 for { select { case <-doneCh } }
	for range doneCh {
		return
	}

}
