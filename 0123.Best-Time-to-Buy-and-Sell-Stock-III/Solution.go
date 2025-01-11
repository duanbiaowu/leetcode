package leetcode

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iii-by-wrnt/
// DP 可以优化为 4 个滚动变量
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// 第一次买入和卖出的利润
	firstBuy, firstSell := -prices[0], 0
	// 第二次买入和卖出的利润
	secondBuy, secondSell := -prices[0], 0

	for i := 1; i < n; i++ {
		firstBuy = max(firstBuy, -prices[i])
		firstSell = max(firstSell, firstBuy+prices[i])

		secondBuy = max(secondBuy, firstSell-prices[i])
		secondSell = max(secondSell, secondBuy+prices[i])
	}

	return secondSell
}

// 状态转移
// dp[i][0]: 第 i 天无任何交易的最大利润。
// dp[i][1]: 第 i 天进行了第一次买入后的最大利润
// dp[i][2]: 第 i 天进行了第一次卖出后的最大利润
// dp[i][3]: 第 i 天进行了第二次买入后的最大利润
// dp[i][4]: 第 i 天进行了第二次卖出后的最大利润

// 状态转移方程如下
// dp[i][1] = max(dp[i-1][1], -prices[i])
// dp[i][2] = max(dp[i-1][2], dp[i-1][1] + prices[i])
// dp[i][3] = max(dp[i-1][3], dp[i-1][2] - prices[i])
// dp[i][4] = max(dp[i-1][4], dp[i-1][3] + prices[i])
func maxProfitDp(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	dp := make([][5]int, n)

	// 初始化状态
	dp[0][1] = -prices[0] // 第一次买入
	dp[0][2] = 0          // 第一次卖出
	dp[0][3] = -prices[0] // 第二次买入
	dp[0][4] = 0          // 第二次卖出

	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][1], -prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])

		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[n-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
