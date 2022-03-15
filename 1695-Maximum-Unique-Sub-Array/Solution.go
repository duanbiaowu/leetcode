package leetcode

import "math"

func maximumUniqueSubarray(nums []int) int {
	maxValue, sum := math.MinInt32, 0
	win := make(map[int]int)
	left, right := 0, 0

	for right < len(nums) {
		sum += nums[right]
		win[nums[right]]++
		for win[nums[right]] > 1 {
			sum -= nums[left]
			win[nums[left]]--
			left++
		}
		maxValue = max(maxValue, sum)
		right++
	}

	if maxValue == math.MinInt32 {
		return 0
	}
	return maxValue
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
