package leetcode

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	// dp[i][j] 表示字符串 S 下标从0到i的子字符串（记为S[0..i]）中
	// 等于字符串 T 下标从0到j的子字符串（记为T[0..j]）的子序列的数目

	// 如果字符串S的长度是m，字符串T的长度是n，
	// 那么 dp[m-1][n-1] 就是字符串S中等于字符串T的子序列的数目

	// Example:
	// S = appplep, T = apple
	// 1. app p le p
	// 2. ap p ple p
	// 3. a p pple p

	// 一个简化代码、提高代码可读性的小技巧为:
	// 将从左向右遍历的方式 改为从右到左遍历
	// 也就是从目标字符串 t 的最后一个字符开始进行匹配
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// 两个字符串中的 “当前字符“ 相等时
				// s[i] 可以选择是否和 t[j] 进行匹配
				// 如果匹配，那么 dp[i][j] 其中一部分数量就是 dp[i+1][j+1]
				// 如果不匹配，那么这样可以让前面的字符跟t[j]匹配，毕竟 t 短 s 长
				//	 dp[i][j] 另外一部分就是 dp[i+1][j]
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

func numDistinct2(s string, t string) int {
	m, n := len(s)+1, len(t)+1
	if m < n {
		return 0
	}

	// dp[i][j] 表示 s[:i]的子序列在 t[:j]出现的次数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

		// 如果 t 为空字符串，那么任意 s[:i] 子序列都能匹配
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		// 如果 s 长度为 0，t 非空时，无子序列匹配
		dp[0][j] = 0
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if s[i-1] == t[j-1] {
				// 如果前一个字符相等
				// 选择匹配当前字符 + 不选择匹配当前字符
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				// 如果前一个字符不相等
				// 直接跳过当前字符
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func numDistinctOpt(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}

	dp := make([]int, len(t)+1)
	dp[0] = 1

	for i := range s {
		for j := min(i, len(t)-1); j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}

	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
