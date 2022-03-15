package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	minLen, sum := math.MaxInt32, 0
	left, right := 0, 0

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
