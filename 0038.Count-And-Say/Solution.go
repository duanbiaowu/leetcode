package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	res := "1"
	var cur strings.Builder
	for i := 2; i <= n; i++ {
		cur.Reset()
		for start, end := 0, 0; end < len(res); start = end {
			for end < len(res) && res[end] == res[start] {
				end++
			}
			cur.WriteString(strconv.Itoa(end - start))
			cur.WriteByte(res[start])
		}
		res = cur.String()
	}
	return res
}
