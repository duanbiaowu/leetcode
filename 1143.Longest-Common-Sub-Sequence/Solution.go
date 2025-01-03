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

// 进一步优化空间: 使用一维滚动数组保存状态转移变化过程中的值
// dp[j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度
// cur 表示 dp[j]，即 dp[i][j]
// prev 表示 dp[j-1]，即 dp[i-1][j-1]
func longestCommonSubsequence2(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	if t1Len < t2Len {
		return longestCommonSubsequence2(text2, text1)
	}

	dp := make([]int, t2Len+1)

	for i := 0; i < t1Len; i++ {
		prev := dp[0]

		for j := 0; j < t2Len; j++ {
			var cur int
			if text1[i] == text2[j] {
				cur = prev + 1
			} else {
				cur = max(dp[j], dp[j+1])
			}

			prev = dp[j+1]
			dp[j+1] = cur
		}
	}

	return dp[t2Len]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
