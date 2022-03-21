package leetcode

// https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	if i < 0 {
		return
	}
	// 1. 从后向前查找第一个顺序对，满足 nums[i] < nums[i+1]
	// 2. 找出另一个最大索引 j 满足 nums[j] > nums[i]
	// 3. 交换 nums[i] 和 nums[j]
	// 4. 最后翻转 nums[i+1:]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 如果在步骤 1 找不到顺序对，说明当前序列已经是一个降序序列，即最大的序列，
	// 直接跳过步骤 2 执行步骤 3，即可得到最小的升序序列
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n>>1; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
