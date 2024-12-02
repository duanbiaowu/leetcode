package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := [26]int{}
	for _, char := range s {
		m[char-'a']++
	}
	for _, char := range t {
		m[char-'a']--
		if m[char-'a'] < 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := [26]int{}

	for _, c := range s {
		cnt[c-'a']++
	}
	for _, c := range t {
		if cnt[c-'a'] == 0 {
			return false
		}
		cnt[c-'a']--
	}

	return true
}
