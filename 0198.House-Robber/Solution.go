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
	dp[1] = nums[1]
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
	second := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func rob3(nums []int) int {
	n := len(nums)
	memo := make([]int, n) // 记忆化: 避免重复计算
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	return robRecursively(nums, 0, &memo)
}

func robRecursively(nums []int, begin int, memo *[]int) int {
	if begin >= len(nums) {
		return 0
	}
	if (*memo)[begin] != -1 {
		return (*memo)[begin]
	}
	res := 0
	for i := begin; i < len(nums); i++ {
		res = max(res, nums[i]+robRecursively(nums, i+2, memo))
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
