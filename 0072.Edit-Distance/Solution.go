package leetcode

// 本质不同的操作实际上只有三种
// 1. 在单词 word1 中插入一个字符
// 2. 在单词 word2 中插入一个字符
// 3. 修改单词 word1 的一个字符

// 状态方程
// 1. dp[i][j] 表示word1的前i个字母转换成word2的前j个字母所使用的最少操作
// 2. word1[i] == word2[j]，dp[i][j] = dp[i-1][j-1]
// 3. word1[i] != word2[j]，dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
