package leetcode

// f(i, j) = max{f(i - 1, j), f(i, j - 1)} + grid[i][j]
// 因为当前状态只和上方的状态和左侧的状态有关，
// 和其它所有状态都无关，所以用完丢弃即可
// 滚动数组: dp[][] 优化为 dp[]
func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	// 滚动数组: 理解为什么长度是 n+1
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[j] = max(dp[j], dp[j-1]) + grid[i-1][j-1]
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
