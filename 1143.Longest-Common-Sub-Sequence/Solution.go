package leetcode

// Example:
// text1 = "abcde", text2 = "ace", answer = 3
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// dp[i][j] 表示 text1 ​ [0:i] 和 text2​ [0:j] 的最长公共子序列的长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
