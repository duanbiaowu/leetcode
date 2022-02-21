package leetcode

import "sort"

func pileBox(box [][]int) int {
	res := 0
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	dp := make([]int, len(box))
	for i := range box {
		// [0]: 长  [1]: 宽  [2]: 高
		dp[i] = box[i][2]
		for j := 0; j < i; j++ {
			if box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
				dp[i] = max(dp[i], dp[j]+box[i][2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
