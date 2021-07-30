package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, min := 0, prices[0]
	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}
	return max
}
