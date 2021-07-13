package leetcode

// 暴力解法(超时)
func lengthOfLongestSubstring(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if !hasDuplicateChar(s, i, j) {
				res = max(res, j-i)
			}
		}
	}
	return res
}

func hasDuplicateChar(s string, start, end int) bool {
	sMap := make(map[uint8]int)

	for i := start; i < end; i++ {
		if _, ok := sMap[s[i]]; ok {
			return true
		}
		sMap[s[i]] = 1
	}

	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
