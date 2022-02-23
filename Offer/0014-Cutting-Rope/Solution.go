package leetcode

// 要求：n>1 并且 m>1
func cuttingRope(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		// j<=i/2: 1*i 和 i*1 是一样的，没必要重复计算
		for j := 1; j <= i>>1; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func cuttingRope2(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	sum := 1
	for n > 4 {
		sum *= 3
		n -= 3
	}
	return sum * n
}
