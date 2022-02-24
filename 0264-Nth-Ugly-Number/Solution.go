package leetcode

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		val := min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		dp[i] = val
		if val%2 == 0 {
			p2++
		}
		if val%3 == 0 {
			p3++
		}
		if val%5 == 0 {
			p5++
		}
	}

	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
