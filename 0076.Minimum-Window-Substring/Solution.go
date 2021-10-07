package leetcode

import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := [128]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	count, start, size := len(t), 0, math.MaxInt32

	for left, right := 0, 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			count--
		}
		need[s[right]]--

		if count == 0 {
			for left < right && need[s[left]] < 0 {
				need[s[left]]++
				left++
			}

			if right-left+1 < size {
				size = right - left + 1
				start = left
			}

			need[s[left]]++
			left++
			count++
		}
	}

	if size == math.MaxInt32 {
		return ""
	}
	return s[start : start+size]
}
