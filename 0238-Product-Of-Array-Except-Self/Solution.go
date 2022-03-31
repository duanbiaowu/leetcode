package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right := 1, 1

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = left
		left *= nums[i]
	}
	for i := n - 1; i > 0; i-- {
		right *= nums[i]
		res[i-1] *= right
	}

	return res
}
