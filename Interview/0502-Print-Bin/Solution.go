package leetcode

import "strings"

func printBin(num float64) string {
	var buf strings.Builder
	buf.WriteString("0.")
	for num != 0 {
		num *= 2
		if num >= 1 {
			buf.WriteByte('1')
			num -= 1
		} else {
			buf.WriteByte('0')
		}
		if buf.Len() > 32 {
			return "ERROR"
		}
	}
	return buf.String()
}
