package leetcode

import "strings"

func fullJustify(words []string, k int) []string {
	var result []string
	var cur []string
	curLen := 0
	for _, word := range words {
		// len(cur) 每个单词一个空格
		if curLen+len(cur)+len(word) > k {
			if len(cur) == 1 {
				result = append(result, cur[0]+strings.Repeat(" ", k-curLen))
			} else {
				spaces := k - curLen
				spacesBetween, extras := spaces/(len(cur)-1), spaces%(len(cur)-1)
				for i := 0; i < extras; i++ {
					cur[i] += " "
				}
				result = append(result, strings.Join(cur, strings.Repeat(" ", spacesBetween)))
			}
			cur = make([]string, 0)
			curLen = 0
		}
		cur = append(cur, word)
		curLen += len(word)
	}

	if len(cur) > 0 {
		result = append(result, strings.Join(cur, " ")+strings.Repeat(" ", k-curLen-len(cur)+1))
	}
	return result
}
