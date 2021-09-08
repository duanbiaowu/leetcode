package leetcode

func wordBreak(s string, wordDict []string) bool {
	set := map[string]struct{}{}
	for i := 0; i < len(wordDict); i++ {
		set[wordDict[i]] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// map[type]type 基本类型时，不存在元素不会引起Panic
			// if dp[j] && set[s[j:i]] {
			if _, ok := set[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
