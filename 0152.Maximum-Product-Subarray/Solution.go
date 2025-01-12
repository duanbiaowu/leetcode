package leetcode

// 求最大值, 可以看做是: 被0拆分的各个子数组的最大值
// 各个子数组最大值计算如下:
//  1. 负数个数为偶数时候，全部相乘最大
//  2. 负数个数为奇数时候，它的左右两边的负数个数一定为偶数，只需求两边最大值
//  3. 遇到0, 重置
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 注意: 最大值可能是由于 2 个以上的负数乘积得出来的
	max, res := 1, nums[0]
	// 计算前缀数组最大值
	for i := 0; i < n; i++ {
		max *= nums[i]
		if max > res {
			res = max
		}
		// 需要规避的数字只有 0
		if nums[i] == 0 {
			max = 1
		}
	}

	max = 1
	// 计算后缀数组最大值
	// 为什么还要计算后缀最大值？
	// 最大值可能存在于最右侧
	// Example: [3,-1,4]
	// 前缀最大值为 3, 显然不符合结果
	// 所以为了处理边界情况下的极端情况，需要再次计算后缀最大值
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

func maxProductDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxF, minF, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 反转很巧妙
		if nums[i] < 0 {
			maxF, minF = minF, maxF
		}
		maxF = max(nums[i], maxF*nums[i])
		minF = min(nums[i], minF*nums[i])
		res = max(res, maxF)
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
