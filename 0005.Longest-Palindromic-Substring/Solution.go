package leetcode

// 暴力搜索(不会超时)
func longestPalindromeBruteForce(s string) string {
	if len(s) <= 1 {
		return s
	}

	res := ""
	maxLen := 0

	for i := 0; i < len(s); i++ {
		// 从当前字符开始，直到字符串结尾
		// 检测可以形成的最大回文字符串长度
		for j := len(s); j > i+maxLen; j-- {
			if isPalindrome(s[i:j]) && j-i > maxLen {
				maxLen = j - i
				res = s[i:j]
			}
		}
	}

	return res
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

// 记忆化搜索
func longestPalindromeMemo(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	res := ""
	maxLen := 0

	// -1: 初始状态
	// 0: 表示 s[i:j] 不是回文串
	// 1: 表示 s[i:j] 是回文串
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := len(s); j > i+maxLen; j-- {
			if isPalindromeMemo(s, i, j-1, memo) > 0 && j-i > maxLen {
				maxLen = j - i
				res = s[i:j]
			}
		}
	}

	return res
}

// 记忆化检测回文字符串
func isPalindromeMemo(s string, i, j int, memo [][]int) int {
	if i == j {
		return 1
	}
	if memo[i][j] != -1 {
		// 代码不会执行到这里
		return memo[i][j]
	}

	// 默认 s[i:j] 不是回文字符串
	memo[i][j] = 0

	if s[i] == s[j] && (j-i <= 1 || isPalindromeMemo(s, i+1, j-1, memo) > 0) {
		memo[i][j] = 1
	}

	return memo[i][j]
}

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

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	left, right := 0, 0
	for i := range s {
		// 返回两个变量（起始位置和结束位置）比单个变量（总长度）可读性更好
		cnt1 := expandAroundCenterCount(s, i, i)   // 单个字符向两边扩展
		cnt2 := expandAroundCenterCount(s, i, i+1) // 相邻字符向两边扩展
		cnt := max(cnt1, cnt2)
		if cnt > right-left {
			left = i - (cnt-1)>>1 // 更新最长子串的左边界
			right = i + cnt>>1    // 更新最长子串的右边界
		}
	}
	return s[left : right+1]
}

func expandAroundCenterCount(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// DP
func longestPalindromeDP(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// 最大回文起始位置
	begin := 0
	// 最大回文长度, 通过起始位置+长度, 可以直接返回结果
	maxLen := 1

	// dp[i][j] 表示 s[i..j] 是否是回文串
	dp := [][]bool{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, n))
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
			// 更新最大回文长度和最大回文起始位置
			if dp[i][j] && j-i+1 > maxLen {
				begin = i
				maxLen = j - i + 1
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
