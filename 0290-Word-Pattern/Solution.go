package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	// 双射
	word2Ch := map[string]byte{}
	ch2Word := map[byte]string{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		ch := pattern[i]
		if word2Ch[word] > 0 && word2Ch[word] != ch ||
			ch2Word[ch] != "" && ch2Word[ch] != word {
			return false
		}
		word2Ch[word] = ch
		ch2Word[ch] = word
	}

	return true
}
