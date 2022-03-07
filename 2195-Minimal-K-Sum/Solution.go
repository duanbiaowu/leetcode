package leetcode

import "sort"

func minimalKSum(nums []int, k int) int64 {
	sum := 0
	// 加入两个哨兵，排序后自动到 nums[0], nums[len - 1]
	nums = append(nums, 0, 2e9)
	sort.Ints(nums)

	for i := 1; ; i++ {
		fill := nums[i] - nums[i-1] - 1
		// 没有可以填充的位置
		if fill <= 0 {
			continue
		}
		if fill >= k {
			// 填充剩余 k 个数：等差数列求和
			// nums[i-1] + nums[i-1] + k + 1
			sum += (nums[i-1]<<1 + k + 1) * k / 2
			return int64(sum)
		}
		// 填充 fill 个数：等差数列求和（+1 和 -1 抵消）
		sum += (nums[i-1] + nums[i]) * fill / 2
		k -= fill
	}
}
