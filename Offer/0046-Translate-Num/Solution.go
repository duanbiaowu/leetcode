package leetcode

import (
	"strconv"
)

// reference: https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/comments/311098
func translateNum(num int) int {
	if num < 0 {
		return 0
	}
	if num <= 9 {
		return 1
	}
	n := num % 100
	// 0-9 或者 >= 26，余数不能按照 2 位数字组合
	// 56，只能拆分为 5 和 6
	// 25，可以拆分为 2 和 5，也可以作为 25
	if n <= 9 || n >= 26 {
		return translateNum(num / 10)
	}
	// 10 - 25，既可以当做一个字母，也可以当做两个字母
	return translateNum(num/10) + translateNum(num/100)
}

func translateNum2(num int) int {
	if num < 0 {
		return 0
	}

	str := strconv.Itoa(num)
	n := len(str)
	// dp[i] 表示从 0-i 位数字的翻译方法数量
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if str[i-2] == '1' || str[i-2] == '2' && str[i-1] <= '5' {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
