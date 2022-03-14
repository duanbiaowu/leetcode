package leetcode

func countSubstrings(s string) int {
	cnt := 0
	for i := range s {
		cnt += count(s, i, i)   // 单个字符向两边扩展
		cnt += count(s, i, i+1) // 相邻字符向两边扩展
	}
	return cnt
}

func count(s string, left, right int) int {
	cnt, n := 0, len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		cnt++
		left--
		right++
	}
	return cnt
}
