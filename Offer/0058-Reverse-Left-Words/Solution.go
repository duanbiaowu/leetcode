package leetcode

func reverseLeftWords(s string, n int) string {
	size := len(s) - 1
	reverseString(&s, 0, size)
	reverseString(&s, 0, size-n)
	reverseString(&s, size-n+1, size)
	return s
}

func reverseString(s *string, left, right int) {
	buf := ([]byte)(*s)
	for left < right {
		buf[left], buf[right] = buf[right], buf[left]
		left++
		right--
	}
	*s = string(buf)
}
