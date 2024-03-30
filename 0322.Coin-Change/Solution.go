package leetcode

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			// 只有硬币金额小于当前金额，才可以拼凑
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
