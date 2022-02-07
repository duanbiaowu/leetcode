package leetcode

import "strings"

func replaceSpaces(S string, length int) string {
	var buf strings.Builder
	space := 0
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			space++
		}
	}
	buf.Grow(space*2 + length)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			buf.WriteByte('%')
			buf.WriteByte('2')
			buf.WriteByte('0')
		} else {
			buf.WriteByte(S[i])
		}
	}

	return buf.String()
}
