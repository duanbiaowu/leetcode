package leetcode

func minSubArrayLen(target int, nums []int) int {
	// 初始化最小长度为数组长度 + 1
	minLen := len(nums) + 1
	sum, left := 0, 0

	for right := range nums {
		sum += nums[right]

		// 找到符合条件的子数组时，开始收缩窗口大小
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	// 数组中不存在符合条件的子数组
	if minLen > len(nums) {
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
