package leetcode

func findAnagrams(s string, p string) []int {
	var res []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return res
	}

	// s 和 p 仅包含小写字母
	win := [26]int{}
	for i := range p {
		win[p[i]-'a']++
	}

	left, right := 0, 0
	for right < sLen {
		index := s[right] - 'a'
		win[index]--

		// 窗口大于 p 的长度
		// 移动左指针，直到窗口小于等于 p 的长度
		for win[index] < 0 {
			win[s[left]-'a']++
			left++
		}

		// 如果当前窗口等于 p 的长度
		// 将当前左指针加入结果集
		if right-left+1 == pLen {
			res = append(res, left)
		}
		right++
	}

	return res
}
