package leetcode

func countPrimeSetBits(left int, right int) int {
	var res int

	for ; left <= right; left++ {
		if isPrime(hammingWeight(uint32(left))) {
			res++
		}
	}

	return res
}

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
