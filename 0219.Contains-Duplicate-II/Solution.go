package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if _, ok := m[v]; ok && i-m[v] <= k {
			return true
		}
		m[v] = i

		// 可选优化项 (滑动窗口)
		// 保持Map长度为k,类似于最小堆的思想
		// 二叉搜索树超时
		// if len(m) > k {
		// 	delete(m, nums[i-k])
		// }
	}

	return false
}
