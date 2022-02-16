package leetcode

func respace(dictionary []string, sentence string) int {
	n := len(sentence)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i] + 1
		for j := range dictionary {
			size := len(dictionary[j])
			if size <= i+1 {
				if dictionary[j] == sentence[i+1-size:i+1] {
					dp[i+1] = min(dp[i+1], dp[i+1-size])
				}
			}
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
