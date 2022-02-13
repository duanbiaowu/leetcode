package leetcode

func waysToChange(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	coins := [4]int{1, 5, 10, 25}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % 1000000007
}
