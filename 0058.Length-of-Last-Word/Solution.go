package leetcode

func lengthOfLastWord(s string) int {
	tail := len(s) - 1
	if tail < 0 {
		return 0
	}

	length := 0

	for s[tail] == ' ' {
		tail--
	}
	for tail >= 0 && s[tail] != ' ' {
		tail--
		length++
	}

	return length
}
