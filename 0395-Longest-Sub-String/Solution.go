package leetcode

import "strings"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/solution/jie-ben-ti-bang-zhu-da-jia-li-jie-di-gui-obla/
func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	cnt := make(map[byte]int)
	for i := range s {
		if _, ok := cnt[s[i]]; ok {
			cnt[s[i]]++
		} else {
			cnt[s[i]] = 1
		}
	}

	for c, v := range cnt {
		if v < k {
			res := 0
			for _, subStr := range strings.Split(s, string(c)) {
				res = max(res, longestSubstring(subStr, k))
			}
			return res
		}
	}

	return len(s)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
