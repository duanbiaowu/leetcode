package leetcode

// 需要注意的是:
// 题目要求, 满足下列条件，就是斐波那契数列，并不是必须按照 1, 2, 3, 5, 8, 13 这样的顺序
// 例如该序列中 [1, 2, 5, 6, 7], 最长序列有 2 中，且长度都是 3
// [1, 6, 7] 和 [2, 5, 7]
// 所以我们可以直接使用 dp[i][j] = dp[j][k] + 1 来计算
func lenLongestFibSubseq(arr []int) int {
	numMap := make(map[int]int)
	for i, v := range arr {
		numMap[v] = i
	}

	// dp[i][j] 表示以 arr[i] 为最后一个数字
	//                arr[j] 为倒数第二个数字
	// 可以形成的最长斐波那契数列的长度
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	// 既然是斐波那契数列，那么最短的长度就是 3
	// s3 = s1 + s2
	res := 2

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if k, ok := numMap[arr[i]-arr[j]]; ok && k < j {
				dp[i][j] = dp[j][k] + 1
			} else {
				dp[i][j] = 2
			}

			res = max(res, dp[i][j])
		}
	}

	// 如果 res > 2 说明找到了长度大于 2 的斐波那契数列
	if res > 2 {
		return res
	}
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
