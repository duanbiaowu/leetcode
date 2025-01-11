package leetcode

// 暴力: 超时，时间复杂度 O(N^2)
// 最大的原因在产生冗余计算
// 例如: 第 1 天价格低于第 2 天价格，即第 1 天成本更低
// 那么我们一定会选择在第 1 天就买入，而不会等到第 2 天
func maxProfitSample(prices []int) int {
	var res int

	n := len(prices)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j] > prices[i] {
				res = max(res, prices[j]-prices[i])
			}
		}
	}

	return res
}

// DP
// 只能选择 某一天 买入这只股票
// 并选择在 未来的某一个不同的日子 卖出该股票
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit, cost := 0, prices[0]
	for _, p := range prices[1:] {
		if p < cost {
			cost = p
		} else if p-cost > profit {
			profit = p - cost
		}
	}

	return profit
}

// 单调栈
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 递增单调栈
	// 其中: 栈底元素是买入价格，栈顶元素是卖出价格
	// 只要保持栈的单调递增性，那么利润就是最高的
	stack := []int{prices[0]}
	res := 0

	for i := 1; i < len(prices); i++ {
		// 如果当前股票价格比栈顶股票价格高
		// 说明当前价格是已知的最高卖出价格
		// 此时可以继续维持递增单调栈
		// 直接将当前价格入栈即可
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			// 维护栈的单调递增性

			// 如果当前股票价格比栈顶股票价格低
			// 此时递增单调栈被破坏
			// 到了该卖出 高于当前股票价格 的股票了
			// 如果不卖出的话，利润就会降低
			index := len(stack) - 1
			for index >= 0 && stack[index] >= prices[i] {
				index--
			}

			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}

		res = max(res, stack[len(stack)-1]-stack[0])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
