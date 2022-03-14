package leetcode

// reference: https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/solution/si-lu-qing-xi-ming-liao-kan-bu-dong-bu-cun-zai-de-/
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	max, min := nums[0], nums[n-1]
	left, right := 0, -1
	for i := 0; i < n; i++ {
		// 从左到右维持最大值，寻找右边界
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}

		// 从右到左维持最小值，寻找左边界
		if nums[n-i-1] > min {
			left = n - i -1
		} else {
			min = nums[n-i-1]
		}
	}

	return right-left+1
}
