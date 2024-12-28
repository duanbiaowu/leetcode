package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	// dp[i][j] 的值代表走到 (i,j) 的最小路径和
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	// 起点路径值
	dp[0][0] = grid[0][0]

	// 计算左边界 (第一列) 单元格的距离
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 计算上边界 (第一行) 单元格的距离
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 计算到达每个单元格的距离
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][cols-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
