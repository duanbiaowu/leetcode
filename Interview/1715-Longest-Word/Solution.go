package leetcode

import (
	"sort"
)

func longestWord(words []string) string {
	// 将单词列表按照长度降序，长度相同按照字典升序
	sort.Slice(words, func(i, j int) bool {
		m, n := len(words[i]), len(words[j])
		if m == n {
			for k := range words[i] {
				if words[i][k] != words[j][k] {
					return words[i][k] < words[j][k]
				}
			}
			return true
		}
		return m > n
	})

	// 把数组转为 Set，提高检测效率
	set := make(map[string]struct{})
	for i := range words {
		set[words[i]] = struct{}{}
	}

	// 遍历检测单词是不是组合单词
	for i := range words {
		// 先对当前字符串剪枝
		delete(set, words[i])
		if isComposeWord(set, words[i]) {
			return words[i]
		}
		set[words[i]] = struct{}{}
	}
	return ""
}

// 检测 word 是否是由 set 里的单词组成
func isComposeWord(set map[string]struct{}, word string) bool {
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(word); i++ {
		if _, ok := set[word[:i+1]]; ok {
			if isComposeWord(set, word[i+1:]) {
				return true
			}
		}
	}
	return false
}
