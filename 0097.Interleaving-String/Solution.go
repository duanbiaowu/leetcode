package leetcode

// https://leetcode-cn.com/problems/interleaving-string/solution/jiao-cuo-zi-fu-chuan-by-leetcode-solution/
func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n == 0 && m == 0 && t == 0 {
		return true
	}
	if n+m != t {
		return false
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	// 空字符串可以交错组成空字符串
	dp[0][0] = true

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 计算当前字符在 s3 中的位置
			p := i + j - 1
			// 检查 s1 的当前字符是否与 s3 的当前字符匹配
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			// 检查 s2 的当前字符是否与 s3 的当前字符匹配
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return dp[n][m]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n == 0 && m == 0 && t == 0 {
		return true
	}
	if n+m != t {
		return false
	}

	dp := make([]bool, m+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || (dp[j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return dp[m]
}
