package leetcode

func strStr(haystack string, needle string) int {
	strLen, subLen := len(haystack), len(needle)

	for fast := 0; fast <= strLen-subLen; fast++ {
		slow := 0
		for ; slow < subLen; slow++ {
			if haystack[fast+slow] != needle[slow] {
				break
			}
		}
		if slow == subLen {
			return fast
		}
	}
	return -1
}
