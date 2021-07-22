package leetcode

import (
	"math"
)

func myAtoi(s string) int {
	left, right := 0, len(s)

	for left < right && s[left] == ' ' {
		left++
	}
	if left == right {
		return 0
	}
	if !isDigit(s[left]) && !isSign(s[left]) {
		return 0
	}

	res := 0
	isNeg := s[left] == '-'
	if !isDigit(s[left]) {
		left++
	}

	// math.MaxIni32: 2147483647
	// math.MinInt32: -2147483648
	for ; left < right && isDigit(s[left]); left++ {
		res = res*10 + int(s[left]-'0')
		if !isNeg && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if isNeg && res > math.MaxInt32 {
			return math.MinInt32
		}
	}
	if isNeg {
		return -res
	}
	return res
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSign(c byte) bool {
	return c == '+' || c == '-'
}
