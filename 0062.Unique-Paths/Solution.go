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

	// 计算 F（i，j）时需要用到 F（i-1，j）和 F（i，j-1）的值
	// 接下来在计算 F（i，j+1）时需要用到 F（i-1，j+1）和 F（i，j）的值
	// 在计算完 F（i，j）之后，就不再需要 F（i-1，j）的值

	// 在二维表格中，F（i，j）和 F（i-1，j）是上下相邻的两个位置
	// 由于在用 F（i-1，j）计算出 F（i，j）之后就不再需要 F（i-1，j）
	// 因此可以只用一个位置来保存 F（i-1，j）和 F（i，j）的值
	// 这个位置在计算 F（i，j）之前保存的是 F（i-1，j）的值
	//        在计算 F（i，j）之后保存的是 F（i，j）的值
	// 由于每个位置能够用来保存两个值，因此只需要一个一维数组就能保存表格中的两行值
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
