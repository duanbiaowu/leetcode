package leetcode

// https://leetcode-cn.com/problems/unique-paths-ii/solution/bu-tong-lu-jing-ii-by-leetcode-solution-2/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j > 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	// 状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 其中 dp[i][j] 表示从左上角走到坐标 (i-1,j-1) 的路径数量
	// 因为索引从 0 开始，为了避免越界，dp 的大小为 m+1 行，n+1 列
	// 所以 dp 的索引要比 obstacleGrid 大 1
	// dp = [m+1][n+1]
	// 最终遍历时，从坐标 (1,1) 开始，所以 dp[1][1] = 1
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[1][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
