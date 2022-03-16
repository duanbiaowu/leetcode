package leetcode

// 变通：题目要求首尾串最大点数，反向过来就是求非首尾串的连续序列的最小点数
// 维护一个 len-k 窗口，保证窗口里面和最小，然后剩余的 k 个数的和就是最大
func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := range cardPoints {
		sum += cardPoints[i]
	}
	cnt, minValue := 0, sum
	win := len(cardPoints) - k
	for i := range cardPoints {
		cnt += cardPoints[i]
		if i >= win {
			cnt -= cardPoints[i-win]
		}
		if i >= win-1 {
			minValue = min(minValue, cnt)
		}
	}

	return sum - minValue
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
