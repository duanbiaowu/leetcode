package leetcode

import (
	"math"
)

// 从底部向顶部迭代
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// 创建一个切片来存储每一层的最小路径和
	// 初始化时多加一个元素以避免边界检查
	dp := make([]int, len(triangle)+1)

	// 从三角形的底部向顶部迭代
	for i := len(triangle) - 1; i >= 0; i-- {
		// 对于当前层的每个元素，计算最小路径和
		for j := 0; j <= i; j++ {
			// 当前元素位置的最小路径和等于:
			// 当前位置的值加上它下面两个位置的最小路径和中的较小值
			// 要从三角形的形态看
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

// 从顶部向底部迭代
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = triangle[0][0]

	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}

	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, dp[i])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
