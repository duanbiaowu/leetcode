package leetcode

// 1. 不要求字典中出现的单词全部都使用
// 2. 字典中的单词可以重复使用
// 如果没有这两个条件，需要 backtrack
func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]struct{})
	for i := range wordDict {
		set[wordDict[i]] = struct{}{}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if _, ok := set[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
