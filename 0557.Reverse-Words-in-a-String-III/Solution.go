package leetcode

import "bytes"

func reverseWords(s string) string {
	res := bytes.Buffer{}
	left, right := 0, len(s)
	for left < right {
		start := left
		for left < right && s[left] != ' ' {
			left++
		}
		for end := left - 1; end >= start; end-- {
			res.WriteByte(s[end])
		}
		for left < right && s[left] == ' ' {
			res.WriteByte(s[left])
			left++
		}
	}

	return res.String()
}
