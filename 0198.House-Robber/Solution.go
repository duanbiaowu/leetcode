package leetcode

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

// 优化DP数组为两个变量: 类似于 Fib数列
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	if n == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(first, nums[1])
	for i := 2; i < n; i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

// 记忆化: 避免重复计算
func robMemo(nums []int) int {
	memo := make([]int, len(nums))

	// 备忘录金额初始化为 -1
	for i := range memo {
		memo[i] = -1
	}

	return robFromMemo(nums, 0, memo)
}

// 参数 memo 已经提前初始化&预分配完成
// 不会再发生扩容操作，所以不需要传递切片指针类型，直接传递切片即可
func robFromMemo(nums []int, begin int, memo []int) int {
	if begin >= len(nums) {
		return 0
	}
	if memo[begin] != -1 {
		return memo[begin]
	}

	// 偷取当前房屋
	// 偷取完当前房屋后，跳过下一个房屋，去偷下下一个房屋
	steal := nums[begin] + robFromMemo(nums, begin+2, memo)

	// 不偷取当前房屋
	// 去偷下一个房屋
	skip := robFromMemo(nums, begin+1, memo)

	memo[begin] = max(steal, skip)

	// 返回两种偷取方案种的最大值
	return memo[begin]
}

// 暴力搜索 (超时)
// 超时原因: 重复检测
func robBruteForce(nums []int) int {
	return robFrom(nums, 0)
}

func robFrom(nums []int, begin int) int {
	if begin >= len(nums) {
		return 0
	}

	// 偷取当前房屋
	// 偷取完当前房屋后，跳过下一个房屋，去偷下下一个房屋
	steal := nums[begin] + robFrom(nums, begin+2)

	// 不偷取当前房屋
	// 去偷下一个房屋
	skip := robFrom(nums, begin+1)

	// 返回两种偷取方案种的最大值
	return max(steal, skip)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
