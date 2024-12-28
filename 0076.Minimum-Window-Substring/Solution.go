package leetcode

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen == 0 || tLen == 0 {
		return ""
	}

	// s 和 t 由英文字母组成
	need := [128]int{}
	for i := 0; i < tLen; i++ {
		need[t[i]]++
	}

	cnt := tLen
	// size 初始化设定为一个不可能的值
	start, size := 0, sLen+1

	left, right := 0, 0
	for right < sLen {
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

	if size == sLen+1 {
		return ""
	}

	return s[start : start+size]
}
