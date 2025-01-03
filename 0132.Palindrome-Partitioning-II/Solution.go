package leetcode

// https://leetcode-cn.com/problems/palindrome-partitioning-ii/solution/xiang-jie-liang-bian-dong-tai-gui-hua-ji-s5xr/
func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	isPal := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPal[i] = make([]bool, n)

		// 单个字符一定是回文串
		for j := 0; j < n; j++ {
			isPal[i][j] = true
		}
	}

	// 从后向前遍历，计算 s[i:j] 是否是回文串
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPal[i][j] = s[i] == s[j] && isPal[i+1][j-1]
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		// 如果 s[0:i] 是回文串，不需要分割
		if isPal[0][i] {
			continue
		}

		// 如果 s[0:i] 不是回文串，就需要分割
		// dp[i] 表示 s[0:i] 的最小分割次数，初始化为一个不可能的值
		dp[i] = n
		for j := 0; j < i; j++ {
			if isPal[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[n-1]
}
