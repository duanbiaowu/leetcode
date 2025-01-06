package leetcode

// 前缀和 + 哈希表
// 前缀和 - 前面的某个前缀和 = 当前区间段的和
// 和当前数组是否已排序无关
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	// 初始化，前缀和为 0 的个数为 1
	// 用于表示 nums[0] == k 的条件分支
	m[0] = 1
	sum, cnt := 0, 0

	for i := range nums {
		sum += nums[i]
		if val, ok := m[sum-k]; ok {
			cnt += val
		}
		m[sum]++
	}

	return cnt
}
