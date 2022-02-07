package leetcode

func CheckPermutation(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	var counter [128]int
	for i := 0; i < n; i++ {
		counter[s1[i]-'0']++
		counter[s2[i]-'0']--
	}
	for i := 0; i < n; i++ {
		if counter[s1[i]-'0'] != 0 {
			return false
		}
	}
	return true
}
