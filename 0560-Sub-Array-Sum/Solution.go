package leetcode

// 前缀和 + 哈希表
func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
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
