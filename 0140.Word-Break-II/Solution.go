package leetcode

import "strings"

func wordBreak(s string, wordDict []string) (sentences []string)  {
	set := map[string]struct{}{}
	for _, w := range wordDict {
		set[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string

	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		var wordsList [][]string
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, ok := set[word]; ok {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}

		word := s[index:]
		if _, ok := set[word]; ok {
			wordsList = append(wordsList, []string{word})
		}

		dp[index] = wordsList
		return wordsList
	}

	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
