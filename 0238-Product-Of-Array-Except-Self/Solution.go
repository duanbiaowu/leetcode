package leetcode

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	// 从左到右计算前缀积
	prefixSum := 1
	for i := range nums {
		res[i] = prefixSum
		prefixSum *= nums[i]
	}

	// 从右到左计算后缀积
	suffixSum := 1
	for i := len(nums) - 1; i > 0; i-- {
		suffixSum *= nums[i]
		res[i-1] *= suffixSum
	}

	return res
}
