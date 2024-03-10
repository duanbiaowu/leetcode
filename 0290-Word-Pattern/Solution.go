package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	// 双射解题
	wordToChar := make(map[string]byte)
	charToWord := make(map[byte]string)

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		char := pattern[i]
		// 如果不满足双射条件，直接返回
		if wordToChar[word] > 0 && wordToChar[word] != char {
			return false
		}
		if len(charToWord[char]) > 0 && charToWord[char] != word {
			return false
		}

		wordToChar[word] = char
		charToWord[char] = word
	}

	return true
}
