package leetcode

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	var buf strings.Builder
	counter := 1
	cur := S[0]
	for i := 1; i < len(S); i++ {
		if S[i] == cur {
			counter++
		} else {
			buf.WriteByte(cur)
			buf.WriteString(strconv.Itoa(counter))
			counter = 1
			cur = S[i]
		}
	}
	buf.WriteByte(cur)
	buf.WriteString(strconv.Itoa(counter))
	if buf.Len() >= len(S) {
		return S
	}
	return buf.String()
}
