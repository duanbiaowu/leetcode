package leetcode

// 暴力求解
func hasTrailingZeros(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (nums[i]|nums[j])&1 == 0 {
				return true
			}
		}
	}

	return false
}

// 优化思路
// 数组中的奇数不需要参与计算过程，因为奇数肯定不尾随 0
// 所以只需要判断偶数的数量是否大于 1 就可以
func hasTrailingZerosOpt(nums []int) bool {
	even := len(nums)
	for _, num := range nums {
		even -= num & 1
	}

	return even > 1
}
