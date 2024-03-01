package leetcode

func maximumStrongPairXor(nums []int) int {
	var res int

	for i := range nums {
		for j := range nums {
			if isStrongPairXor(nums[i], nums[j]) {
				res = max(res, nums[i]^nums[j])
			}
		}
	}

	return res
}

func isStrongPairXor(x, y int) bool {
	diff := x - y
	if diff < 0 {
		diff = -diff
	}
	if x < y {
		return diff <= x
	}
	return diff <= y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
