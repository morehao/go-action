package main

import "fmt"

func main() {
	fmt.Println(countPrimes(100))
}

// 超时警告
func countPrimes1(n int) (cnt int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) (cnt int) {
	isPrimeList := make([]bool, n)
	for i := range isPrimeList {
		isPrimeList[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrimeList[i] {
			cnt++
			// 如果i是质数，2*i则一定不是质数，把不是质数的标记出来，剩下的都是质数
			for j := 2 * i; j < n; j += i {
				isPrimeList[j] = false
			}
		}
	}
	return
}
