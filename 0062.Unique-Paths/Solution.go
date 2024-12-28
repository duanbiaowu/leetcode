package leetcode

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	// dp[i][j] -> 从 (0,0) 到 (i,j) 的路径数目
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		// 第一列，不能继续向左走了
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		// 第一行，不能继续向上走了
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	// 空间优化：每次只需要 dp[i-1][j], dp[i][j-1]
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
