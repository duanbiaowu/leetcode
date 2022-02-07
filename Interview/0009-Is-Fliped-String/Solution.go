package leetcode

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var buf strings.Builder
	buf.WriteString(s2)
	buf.WriteString(s2)
	return strings.Contains(buf.String(), s1)
}
