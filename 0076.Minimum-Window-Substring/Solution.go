package leetcode

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m == 0 || n == 0 {
		return ""
	}

	need := [128]int{}
	for i := 0; i < n; i++ {
		need[t[i]]++
	}

	cnt := n
	// m 设定为一个不可能的值
	start, size := 0, m+1
	left, right := 0, 0

	for right < m {
		if need[s[right]] > 0 {
			cnt--
		}
		need[s[right]]--

		if cnt == 0 {
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
			cnt++
		}

		right++
	}

	if size == m+1 {
		return ""
	}
	return s[start : start+size]
}
