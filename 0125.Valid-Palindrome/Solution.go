package leetcode

import "strings"

func isPalindrome(s string) bool {
	// 只考虑字母和数字字符, 可以忽略字母的大小写
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isChar(s[left]) {
			left++
		}
		for left < right && !isChar(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
