package leetcode

// 尽可能地完成更多的交易（多次买卖一支股票）
// 不能同时参与多笔交易（必须在再次购买前出售掉之前的股票）
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润
// dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润
func maxProfitDP(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	// 全部交易结束后，持有股票的收益一定低于不持有股票的收益
	return dp[n-1][0]
}

// 每一天的状态只与前一天的状态有关，而与更早的状态都无关
// 因此不必存储这些无关的状态
func maxProfitDP2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
