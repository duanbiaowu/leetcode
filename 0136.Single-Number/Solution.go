package leetcode

import "sort"

// 利用异或运算的性质:
// 1. 任何数和 0 做异或运算，结果仍然是原来的数
// 2. 任何数和其自身做异或运算，结果是 0
// 3. 异或运算满足交换律和结合律，a^b^a = b^a^a = b^(a^a) = b^0 = b
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func singleNumberSimple(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	// 排序完成后
	// 2 个数字为 1 对
	// 如果 1 对中 2 个数字不相同，必然是第 1 个数字
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}
