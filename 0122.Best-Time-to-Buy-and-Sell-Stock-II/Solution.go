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

// dp[i][0] 表示第 i 天交易完后卖出股票的最大利润
// dp[i][1] 表示第 i 天交易完后持有 1 支股票的最大利润
func maxProfitDP(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)
	// 第一天买入股票，肯定是亏钱的，负利润
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// 可以在同一天买入并卖出
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 例如 [7,1,5,3,6,4]
		// 那么第 1 天卖出的最大利润应该是 1, 也就是说，第 0 天的时候，不买入价格为 7 的股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	// 全部交易结束后，持有股票的收益一定低于全部卖出股票的收益
	return dp[n-1][0]
}

// 每一天的状态只与前一天的状态有关，而与更早的状态都无关
// 因此不必存储这些无关的状态

// 进一步的优化:
// profit 和 cost 两个变量，只有都为正数时才有影响结果
// 也就是说，只有 买入成本 - 卖出收益 > 0 的时候才有必要进行操作
// 进而可以优化到只有一个返回值变量 profit
// func maxProfit(prices []int) int {
//     profit := 0
//     for i := 1; i < len(prices); i++ {
//		   // 只要当天价格比前一天价格高，就说明存在利润
//		   // 可以在前一天买入，当天买入即可
//         if prices[i] > prices[i-1] {
//             profit += prices[i] - prices[i-1]
//         }
//     }

//     return profit
// }

func maxProfitDP2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// profit: 表示第 i 天卖出股票的利润
	// cost:   表示第 i 天买入股票的利润
	profit, cost := 0, -prices[0]
	for i := 1; i < n; i++ {
		profit = max(profit, cost+prices[i])
		cost = max(cost, profit-prices[i])
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
