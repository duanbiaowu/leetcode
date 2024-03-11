package leetcode

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	slow, fast := 0, 0

	for slow < sLen && fast < tLen {
		if s[slow] == t[fast] {
			slow++
		}
		fast++
	}
	return slow == sLen
}
