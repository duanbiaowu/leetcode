package leetcode

// 求最大值, 可以看做是: 被0拆分的各个子数组的最大值
// 各个子数组最大值计算如下:
// 	1. 负数个数为偶数时候，全部相乘最大
//	2. 负数个数为奇数时候，它的左右两边的负数个数一定为偶数，只需求两边最大值
//  3. 遇到0, 重置
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	max, res := 1, nums[0]
	// 计算前缀数组最大值
	for i := 0; i < n; i++ {
		max *= nums[i]
		if max > res {
			res = max
		}
		if nums[i] == 0 {
			max = 1
		}
	}

	max = 1
	// 计算后缀数组最大值
	for i := n - 1; i >= 0; i-- {
		max *= nums[i]
		if max > res {
			res = max
		}
		if nums[i] == 0 {
			max = 1
		}
	}

	return res
}
