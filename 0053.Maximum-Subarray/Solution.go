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
		// update max value of nums[0] - nums[i]
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
