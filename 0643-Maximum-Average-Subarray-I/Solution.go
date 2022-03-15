package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	if k == 0 {
		return 0
	}
	sum, maxAvg := 0, math.MinInt32
	left, right := 0, 0

	for right < len(nums) {
		sum += nums[right]
		if right-left+1 == k {
			maxAvg = max(maxAvg, sum)
		}
		if right >= k-1 {
			sum -= nums[left]
			left++
		}
		right++
	}

	return float64(maxAvg) / float64(k)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
