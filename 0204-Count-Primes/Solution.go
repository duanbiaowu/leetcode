package leetcode

// 线性筛
func countPrimes(n int) int {
	var primes []int
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for j := range primes {
			if i*primes[j] >= n {
				break
			}
			isPrime[i*primes[j]] = false
			if i%primes[j] == 0 {
				break
			}
		}
	}
	return len(primes)
}

// 埃氏筛
func countPrimes2(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	// 从 2 开始枚举到 sqrt(n)
	for i := 2; i*i < n; i++ {
		// 如果当前是素数
		if isPrime[i] {
			// 就把从 i*i 开始，i 的所有倍数都设置为 false
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	cnt := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
		}
	}
	return cnt
}
