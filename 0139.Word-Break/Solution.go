package leetcode

// 1. 不要求字典中出现的单词全部都使用
// 2. 字典中的单词可以重复使用
// 如果没有这两个条件，需要 backtrack
func wordBreak(s string, wordDict []string) bool {
	// 首先构建字典的 Set, 便于快速查找
	set := make(map[string]struct{})
	for _, word := range wordDict {
		set[word] = struct{}{}
	}

	// 初始化动态规划状态转移数组
	// 状态表达式: dp[i] = dp[j] && check(s[j : i])
	// 其中，dp[i] 表示字符串 s[0 : i-1] 是否存在于字典中
	// check(s[j : i]) 表示子字符串 s[j : i] 是否存在于字典中
	n := len(s)
	dp := make([]bool, n+1)

	// 空字符串可以直接构建
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
