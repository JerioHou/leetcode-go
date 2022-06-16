package leetcode_go

import (
	"testing"
)

/*
    统计所有小于非负整数 n 的质数的数量。
	https://leetcode-cn.com/problems/count-primes/
*/
func TestCountPrimes(t *testing.T) {

	t.Log(countPrimes(100))
}

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	count := 0
	for j := 2; j < n; j++ {
		if isPrime[j] {
			count++
		}
		for k := 2 * j; k < n; k += j {
			isPrime[k] = false
		}
	}
	return count
}
