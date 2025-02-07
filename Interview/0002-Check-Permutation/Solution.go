package leetcode

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	cnt := [128]int{}

	for i := range s1 {
		cnt[s1[i]-'0']++
		cnt[s2[i]-'0']--
	}

	for i := range s1 {
		if cnt[s1[i]-'0'] != 0 {
			return false
		}
	}

	return true
}
