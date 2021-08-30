package leetcode

func trailingZeroes(n int) int {
	if n/5 == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}

func trailingZeroes2(n int) int {
	count := 0
	for n >= 5 {
		count += n / 5
		n /= 5
	}
	return count
}
