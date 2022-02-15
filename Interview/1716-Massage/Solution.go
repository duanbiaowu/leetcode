package leetcode

func massage(nums []int) int {
	a, b := 0, 0
	for i := range nums {
		c := max(b, a+nums[i])
		a = b
		b = c
	}
	return b
}

func massage2(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
