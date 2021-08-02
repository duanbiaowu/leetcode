package leetcode

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}

	win := [26]int{}
	for _, char := range s1 {
		win[char-'a']--
	}

	for left, right := 0, 0; right < m; right++ {
		char := s2[right] - 'a'
		win[char]++
		for win[char] > 0 {
			win[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}

	return false
}
