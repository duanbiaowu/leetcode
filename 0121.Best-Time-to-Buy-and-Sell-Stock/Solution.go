package leetcode

// 只能选择 某一天 买入这只股票
// 并选择在 未来的某一个不同的日子 卖出该股票
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

// 单调栈
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	stack := []int{prices[0]}
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0 && stack[index] >= prices[i]; index-- {

			}

			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		if max < stack[len(stack)-1]-stack[0] {
			max = stack[len(stack)-1] - stack[0]
		}
	}

	return max
}
