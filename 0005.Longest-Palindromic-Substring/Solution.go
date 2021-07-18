package leetcode

// 暴力解法(不会超时)
//func longestPalindrome(s string) string {
//	if len(s) <= 1 {
//		return s
//	}
//
//	res := ""
//	maxLength := 0
//
//	for i := 0; i < len(s); i++ {
//		for j := len(s); j > i+maxLength; j-- {
//			if isPalindrome(s[i:j]) && j-i > maxLength {
//				maxLength = j-i
//				res = s[i:j]
//			}
//		}
//	}
//
//	return res
//}

// 中心扩展(利用回文串性质)
//func longestPalindrome(s string) string {
//	if len(s) <= 1 {
//		return s
//	}
//
//	start, end := 0, 0
//	for i := 0; i < len(s); i++ {
//		left1, right1 := expandAroundCenter(s, i, i)	// 单个字符向两边扩展
//		left2, right2 := expandAroundCenter(s, i, i+1)	// 相邻字符向两边扩展
//		if right1-left1 > end-start {
//			start, end = left1, right1
//		}
//		if right2-left2 > end-start {
//			start, end = left2, right2
//		}
//	}
//
//	return s[start : end+1]
//}

// DP
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 回文起始位置
	begin := 0
	// 回文长度, 通过起始位置+长度, 可以直接返回结果
	maxLen := 1
	n := len(s)

	// dp[i][j] 表示 s[i..j] 是否是回文串
	dp := [][]bool{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, n+1))
	}

	// 初始化：所有长度为 1 的子串都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 递推开始
	// 先枚举子串长度
	for left := 2; left <= n; left++ {
		// 枚举左边界
		for i := 0; i < n; i++ {
			// 由 left 和 i 可以确定右边界，即 j - i + 1 = left
			j := left + i - 1
			if j >= n {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				// 边界情况, "aa" "aba" ...
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文
			// 更新回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n>>1; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
