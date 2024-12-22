package leetcode

// https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/
func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// 1. 从后向前查找第一个顺序对，满足 nums[left] >= nums[left+1]
	// 2. 从后向前查找找出一个最大索引 right 满足 nums[right] > nums[left]
	// 3. 交换 nums[left] 和 nums[right]
	// 4. 最后翻转 nums[left+1:], 也就是从 left+1 到最后的元素，使其变为升序

	// 示例参数: [2,3,1]
	// 因为从后向前的顺序对为 [1,3], 步骤 1 执行完时 left = 0
	left := n - 2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}

	// 如果在步骤 1 找不到顺序对，说明当前序列已经是一个降序序列，即最大的序列
	// 直接跳过步骤 2, 3 执行步骤 4，即可得到最小的升序序列
	if left >= 0 {
		// 如果找到顺序对
		// 此时 left = 0, right = 2
		right := n - 1
		// 从后向前查找找出一个最大索引 right
		// 步骤 2 执行完时 right = 1
		for right >= 0 && nums[left] >= nums[right] {
			right--
		}
		// 交换后的数组变为 [3,2,1]
		nums[left], nums[right] = nums[right], nums[left]
	}

	// 翻转 [2, 1] 变为 [1, 2]
	// 最终返回 [3, 1, 2]
	reverse(nums[left+1:])
}

func reverse(a []int) {
	left, right := 0, len(a)-1

	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
}
