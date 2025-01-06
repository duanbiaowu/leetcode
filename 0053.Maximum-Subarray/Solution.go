package leetcode

// simple solution
//func maxSubArray(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	res, sum := nums[0], nums[0]
//	for i := 1; i < len(nums); i++ {
//		if sum > 0 {
//			sum += nums[i]
//		} else {
//			sum = nums[i]
//		}
//		if sum > res {
//			res = sum
//		}
//	}
//	return res
//}

// dp solution
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		// 类似前缀和的思路
		// 对于当前元素，存在 2 种策略
		// 1. 如果前一个元素大于 0，更新当前值，形成最大子数组和的 “连续滚动“
		// 1. 如果前一个元素小于等于 0，此时可以直接将当前元素直接作为新的 最大子数组的开始，所以不做任何更新
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		// 更新最大子数组和
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
