package leetcode

func maxSubArraySumCircular(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := nums[0]
	preMax, curMax := sum, sum
	preMin, curMin := sum, sum

	for i := 1; i < n; i++ {
		preMax = max(preMax+nums[i], nums[i])
		curMax = max(curMax, preMax)
		preMin = min(preMin+nums[i], nums[i])
		curMin = min(curMin, preMin)
		sum += nums[i]
	}

	if curMax < 0 {
		return curMax
	}
	return max(curMax, sum-curMin)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
