package leetcode

func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, c := range s {
		if cnt[c-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func firstUniqChar2(s string) byte {
	m := make(map[int32]bool)
	for _, c := range s {
		_, ok := m[c]
		m[c] = !ok
	}
	for i, c := range s {
		if m[c] {
			return s[i]
		}
	}
	return ' '
}
