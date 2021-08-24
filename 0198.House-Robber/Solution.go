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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
