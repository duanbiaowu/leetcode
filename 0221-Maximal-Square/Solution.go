package leetcode

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	// dp[i][j]: 以 i 行 j 列为右下角所能构成的最大正方形边长
	// dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	maxSide := 0

	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				if maxSide < dp[i][j] {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

func min(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < z {
		return y
	}
	return z
}
