package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
