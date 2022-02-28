package leetcode

import (
	"fmt"
	"strconv"
)

// reference: https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/solution/jian-zhi-offer-60-n-ge-tou-zi-de-dian-sh-z36d/
func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := range dp {
		dp[i] = float64(1) / float64(6)
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(dp); j++ {
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / float64(6)
			}
		}
		dp = tmp
	}

	for i := range dp {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", dp[i]), 64)
		dp[i] = v
	}
	return dp
}
