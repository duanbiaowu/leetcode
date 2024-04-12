package leetcode

import "fmt"

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

// 说明：
// 使用单纯的暴力法无法解题
// 因为会遇到边界情况 (边界情况建模)
// 边界情况测试用例
// 需要转换到回溯法
func lengthOfLISBruteForce(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 0
	for i := 0; i < n; i++ {
		// 当前
		curNum := nums[i]
		curLen := 1

		fmt.Println(curNum, curLen)

		for j := i + 1; j < n; j++ {
			if nums[j] > curNum {
				curNum = nums[j]
				curLen++
			}
		}

		fmt.Println(curNum, curLen)
		fmt.Println()

		res = max(res, curLen)
	}

	return res
}

func lengthOfLISBruteForce2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 0
	for i := n - 1; i >= 0; i-- {
		// 当前
		curNum := nums[i]
		curLen := 1

		fmt.Println(curNum, curLen)

		for j := i - 1; j >= 0; j-- {
			if curNum > nums[j] {
				curNum = nums[j]
				curLen++
			}
			fmt.Println(curNum, curLen)
		}

		fmt.Println(curNum, curLen)
		fmt.Println()

		res = max(res, curLen)
	}

	return res
}

// 试图转换成对应的迭代版本代码
func lengthOfLIS3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		maxLen = max(maxLen, findLIS(nums, i, nums[i], 1))
	}
	return maxLen
}

func findLIS(nums []int, begin, prev, curLen int) int {
	if begin == len(nums) {
		return curLen
	}
	includedLen := curLen
	if nums[begin] > prev {
		includedLen = findLIS(nums, begin+1, nums[begin], curLen+1)
	}
	notIncludedLen := findLIS(nums, begin+1, prev, curLen)
	return max(includedLen, notIncludedLen)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
