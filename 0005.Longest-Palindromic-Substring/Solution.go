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
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
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
