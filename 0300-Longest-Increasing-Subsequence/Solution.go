package leetcode

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 0
	// dp[i] 的值代表以 nums[i] 结尾的最长子序列长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// reference: https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-2/
func lengthOfLIS2(nums []int) int {
	// tails[k] 的值代表 长度为 k+1 子序列 的尾部元素值
	tails := make([]int, len(nums))
	res := 0

	for i := range nums {
		low, hi := 0, res
		for low < hi {
			mid := low + (hi-low)>>1
			if tails[mid] < nums[i] {
				low = mid + 1
			} else {
				hi = mid
			}
		}
		tails[low] = nums[i]
		if hi == res {
			res++
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
