package leetcode

// 原地哈希
// 每个数字 n 都放到下标为 n-1 的位置上
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 将所有小于等于 0 的元素替换为 n+1
	// [1,2,0] => [1,2,4]
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 将所有小于等于 n 的元素对应的位置标记为负数
	// [1,2,4] => [-1,-2,4]
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	// [-1,-2,4] => return 3
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	// 没有缺失的正整数，返回 n+1
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
