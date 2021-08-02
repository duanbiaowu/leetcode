package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, char := range magazine {
		m[char-'a']++
	}
	for _, char := range ransomNote {
		m[char-'a']--
		if m[char-'a'] < 0 {
			return false
		}
	}

	return true
}
