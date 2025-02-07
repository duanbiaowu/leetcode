package leetcode

import "strings"

func replaceSpaces(S string, length int) string {
	var buf strings.Builder

	spaces := 0
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			spaces++
		}
	}

	buf.Grow(spaces<<1 + length)

	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(S[i])
		}
	}

	return buf.String()
}
