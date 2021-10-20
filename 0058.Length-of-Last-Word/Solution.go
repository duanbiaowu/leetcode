package leetcode

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	if index < 0 {
		return 0
	}

	res := 0
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		res++
		index--
	}
	return res
}
