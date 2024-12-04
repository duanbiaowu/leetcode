package leetcode

func strStr(haystack string, needle string) int {
	strLen, subLen := len(haystack), len(needle)

	for head := 0; head <= strLen-subLen; head++ {
		tail := 0
		for ; tail < subLen; tail++ {
			if haystack[head+tail] != needle[tail] {
				break
			}
		}
		if tail == subLen {
			return head
		}
	}
	return -1
}
