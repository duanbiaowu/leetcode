package leetcode

func findAnagrams(s string, p string) []int {
	var res []int
	m, n := len(s), len(p)
	if m < n {
		return res
	}

	win := [128]int{}
	for i := range p {
		win[p[i]-'a']++
	}

	left, right := 0, 0
	for right < m {
		win[s[right]-'a']--
		for win[s[right]-'a'] < 0 {
			win[s[left]-'a']++
			left++
		}
		if right-left+1 == n {
			res = append(res, left)
		}
		right++
	}

	return res
}
