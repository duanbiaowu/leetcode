package leetcode

func canPartition(nums []int) bool {
	// 要求数组分割为 2 个子集
	// 所以长度不够时直接返回
	if len(nums) < 2 {
		return false
	}

	// 计算数组所有元素的总和 和 数组中的最大元素
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	// 如果数组所有元素的总和为奇数
	// 不可能划分为 2 个相等的子集
	// 直接返回 false
	if sum&1 == 1 {
		return false
	}

	target := sum / 2
	// 如果数组中的最大元素超过数组总和的一半
	// 不可能划分为 2 个相等的子集
	// 直接返回 false
	if max > target {
		return false
	}

	// 状态转移方程: dp[j] = dp[j] || dp[j−nums[i]]
	// 其中 dp[j] 表示当前数字是否可以满足分割 2 个子集
	dp := make([]bool, target+1)
	dp[0] = true

	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}

	return dp[target]
}
